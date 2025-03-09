package shortreq

import (
	"api-shortener/restapi"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RequestCreationError struct {
	err error
}

func (e *RequestCreationError) Error() string {
	return fmt.Sprintf("Error while creating outgoing request: %s", e.err.Error())
}

type IIncomingRequestProcessor interface {
	CreateOutgoingRequest(api *restapi.ShortenedAPI) (*http.Request, error)
}

type IncomingRequestProcessor struct {
	configDAO restapi.IOutgoingRequestConfigDAO
}

func (processor *IncomingRequestProcessor) CreateOutgoingRequest(api *restapi.ShortenedAPI) (*http.Request, error) {
	requestConfigs, _ := processor.configDAO.GetAllByAPIID(api.ID)
	requestConfig := requestConfigs[0]
	request, err := http.NewRequest(requestConfig.Method, requestConfig.Url, strings.NewReader(requestConfig.Body))
	if err != nil {
		return nil, &RequestCreationError{err: err}
	}

	for _, header := range requestConfig.Headers {
		request.Header.Add(header.Name, header.Value)
	}
	q := request.URL.Query()
	for _, param := range requestConfig.Params {
		q.Add(param.Name, param.Value)
	}
	request.URL.RawQuery = q.Encode()
	return request, err
}

func NewIncomingRequestProcessor(configDAO restapi.IOutgoingRequestConfigDAO) IIncomingRequestProcessor {
	return &IncomingRequestProcessor{configDAO: configDAO}
}

type IOutgoingRequestProcessor interface {
	Process(request *http.Request, c *gin.Context, api *restapi.ShortenedAPI)
}

type OutgoingRequestProcessor struct {
	jsonResponseShortener *JSONResponseShortener
	rulesResolver         IRulesResolver
	client                IOutgoingRequestClient
}

func (processor *OutgoingRequestProcessor) Process(request *http.Request, c *gin.Context, api *restapi.ShortenedAPI) {
	response, err := processor.client.MakeRequest(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logrus.Errorf("Error while making request with API %d: %s", api.ID, err.Error())
		return
	}

	body, err := io.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logrus.Errorf("Error decoding response body for API %d: %s", api.ID, err.Error())
		return
	}

	for headerName := range response.Header {
		c.Writer.Header().Add(headerName, response.Header.Get(headerName))
	}

	rules, err := processor.rulesResolver.GetRules(api)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logrus.Errorf("Getting transformation rules for API %d: %s", api.ID, err.Error())
		return
	}

	result, err := processor.jsonResponseShortener.Shorten(body, rules)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logrus.Errorf("Getting shortening response for API %d: %s", api.ID, err.Error())
		return
	}

	c.JSON(response.StatusCode, result)
}

func NewOutgoingRequestProcessor(
	jsonShortener *JSONResponseShortener,
	rulesResolver IRulesResolver,
	client IOutgoingRequestClient,
) IOutgoingRequestProcessor {
	return &OutgoingRequestProcessor{
		jsonResponseShortener: jsonShortener,
		rulesResolver:         rulesResolver,
		client:                client,
	}
}
