package http

import (
	"api-shortener/shortreq"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/oj"
	"github.com/sirupsen/logrus"
)

var (
	errRequestIsAlreadySent          = errors.New("request is already sent")
	errWhileMakingRequest            = errors.New("error while making request to the server")
	errWhileCreatingRequestObject    = errors.New("error while creating request object")
	errWhileReadingServerResponse    = errors.New("error while reading server response")
	errWhileShorteningServerResponse = errors.New("error while shortening server response")
)

type ResponseShorteningService struct {
	configDAO  shortreq.IOutgoingRequestConfigDAO
	headersDAO shortreq.IOutgoingRequestHeaderDAO
	paramsDAO  shortreq.IOutgoingRequestParamDAO
	client     IOutgoingRequestClient
	limiter    ILoopLimiter
}

func (s *ResponseShorteningService) ProcessRequest(api *shortreq.ShortenedAPI) (*ShortenedResponse, error) {
	if !s.limiter.AddNewRequest(api.ID) {
		logrus.Warningf("Max requests limit exceeded for API %d", api.ID)
		return nil, errRequestIsAlreadySent
	}
	defer s.limiter.RemoveRequest(api.ID)

	request, err := s.createOutgoingRequest(api)
	if err != nil {
		logrus.Errorf("Error while creating request for API %d: %s", api.ID, err.Error())
		return nil, err
	}
	return s.processRequest(request, api)
}

func (s *ResponseShorteningService) createOutgoingRequest(api *shortreq.ShortenedAPI) (*http.Request, error) {
	requestConfig, err := s.configDAO.GetByAPIID(api.ID)
	if err != nil {
		return nil, errWhileCreatingRequestObject
	}
	request, err := http.NewRequest(requestConfig.Method, requestConfig.Url, strings.NewReader(requestConfig.Body))
	if err != nil {
		return nil, errWhileCreatingRequestObject
	}

	headers, err := s.headersDAO.GetAllByConfigID(requestConfig.ID)
	if err != nil {
		return nil, errWhileCreatingRequestObject
	}
	for _, header := range headers {
		request.Header.Add(header.Name, header.Value)
	}
	params, err := s.paramsDAO.GetAllByConfigID(requestConfig.ID)
	if err != nil {
		return nil, errWhileCreatingRequestObject
	}
	q := request.URL.Query()
	for _, param := range params {
		q.Add(param.Name, param.Value)
	}
	request.URL.RawQuery = q.Encode()
	return request, nil
}

func (s *ResponseShorteningService) processRequest(request *http.Request, api *shortreq.ShortenedAPI) (*ShortenedResponse, error) {
	response, err := s.client.MakeRequest(request)
	if err != nil {
		logrus.Errorf("Error while making request with API %d: %s", api.ID, err.Error())
		return nil, errWhileMakingRequest
	}

	body, err := io.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		logrus.Errorf("Error decoding response body for API %d: %s", api.ID, err.Error())
		return nil, errWhileReadingServerResponse
	}

	resultHeader := make(http.Header)
	for headerName := range response.Header {
		resultHeader.Add(headerName, response.Header.Get(headerName))
	}

	rules := s.getRules(api)
	result, err := s.shorten(body, rules)
	if err != nil {
		logrus.Errorf("Getting shortening response for API %d: %s", api.ID, err.Error())
		return nil, err
	}

	return &ShortenedResponse{json: &result, statusCode: response.StatusCode, headers: resultHeader}, nil
}

func (processor *ResponseShorteningService) getRules(api *shortreq.ShortenedAPI) map[string]string {
	resultRules := make(map[string]string)
	for _, rule := range api.ShorteningRules {
		resultRules[rule.FieldName] = rule.FieldValueQuery
	}
	return resultRules
}

func (shortener *ResponseShorteningService) shorten(body []byte, rules map[string]string) (map[string]any, error) {
	parsedJson, err := oj.Parse(body)
	if err != nil {
		return map[string]any{}, errWhileShorteningServerResponse
	}

	result, err := shortener.shortenWithRules(parsedJson, rules)
	if err != nil {
		return map[string]any{}, errWhileShorteningServerResponse
	}

	return result, nil
}

func (shortener *ResponseShorteningService) shortenWithRules(json any, rules map[string]string) (map[string]any, error) {
	result := make(map[string]any)
	for rule_k, rule_v := range rules {
		expr, err := jp.ParseString(rule_v)
		if err != nil {
			return map[string]any{}, errWhileShorteningServerResponse
		}
		parsed := expr.Get(json)
		result[rule_k] = parsed
	}
	return result, nil
}

func NewResponseShorteningService(
	configDAO shortreq.IOutgoingRequestConfigDAO,
	headersDAO shortreq.IOutgoingRequestHeaderDAO,
	paramsDAO shortreq.IOutgoingRequestParamDAO,
	client IOutgoingRequestClient,
	limiter ILoopLimiter,
) *ResponseShorteningService {
	return &ResponseShorteningService{
		configDAO:  configDAO,
		headersDAO: headersDAO,
		paramsDAO:  paramsDAO,
		client:     client,
		limiter:    limiter,
	}
}
