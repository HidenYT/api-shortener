package shortreq

import (
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
	CreateOutgoingRequest(api *ShortenedAPI) (*http.Request, error)
}

type IncomingRequestProcessor struct {
	outgoingRequestConfigResolver IOutgoingRequestConfigResolver
}

func (processor *IncomingRequestProcessor) CreateOutgoingRequest(api *ShortenedAPI) (*http.Request, error) {
	requestConfig := processor.outgoingRequestConfigResolver.GetRequestConfigModel(api)

	request, err := http.NewRequest(requestConfig.Method, requestConfig.Url, strings.NewReader(requestConfig.Body))
	if err != nil {
		return nil, &RequestCreationError{err: err}
	}

	for k, v := range requestConfig.Headers {
		for _, val := range v {
			request.Header.Add(k, val)
		}
	}
	q := request.URL.Query()
	for k, v := range requestConfig.Params {
		q.Add(k, v)
	}
	request.URL.RawQuery = q.Encode()
	return request, err
}

func NewIncomingRequestProcessor(outgoingRequestResolver IOutgoingRequestConfigResolver) IIncomingRequestProcessor {
	return &IncomingRequestProcessor{
		outgoingRequestConfigResolver: outgoingRequestResolver,
	}
}

type IOutgoingRequestProcessor interface {
	Process(request *http.Request, c *gin.Context, api *ShortenedAPI)
}

type OutgoingRequestProcessor struct {
	jsonResponseShortener *JSONResponseShortener
	rulesResolver         IRulesResolver
	client                IOutgoingRequestClient
}

func (processor *OutgoingRequestProcessor) Process(request *http.Request, c *gin.Context, api *ShortenedAPI) {
	response, err := processor.client.MakeRequest(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logrus.Errorf("Error while making request with API %d: %s", api.ID, err.Error())
		return
	}

	body, err := io.ReadAll(response.Body)
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
