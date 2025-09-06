package http_shortener

import (
	"errors"
	"net/http"
	"strings"

	shortener "github.com/HidenYT/api-shortener/internal/response-shortener"
	"github.com/HidenYT/api-shortener/internal/shortreq"

	"github.com/sirupsen/logrus"
)

var (
	errRequestIsAlreadySent       = errors.New("request is already sent")
	errWhileCreatingRequestObject = errors.New("error while creating request object")
)

type ResponseShorteningService struct {
	configDAO  shortreq.IOutgoingRequestConfigDAO
	headersDAO shortreq.IOutgoingRequestHeaderDAO
	paramsDAO  shortreq.IOutgoingRequestParamDAO
	shortener  shortener.IResponseShortener
	limiter    ILoopLimiter
}

func (s *ResponseShorteningService) ProcessRequest(api *shortreq.ShortenedAPI) (*shortener.ShortenedResponse, error) {
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

func (s *ResponseShorteningService) processRequest(request *http.Request, api *shortreq.ShortenedAPI) (*shortener.ShortenedResponse, error) {
	result, err := s.shortener.ProcessRequest(request, s.getRules(api))
	if err != nil {
		if errors.Is(err, shortener.ErrWhileMakingRequest) {
			logrus.Errorf("Error while making request to the target server: %s", err.Error())
		} else if errors.Is(err, shortener.ErrWhileReadingServerResponse) {
			logrus.Errorf("Error while reading response from the target server: %s", err.Error())
		} else if errors.Is(err, shortener.ErrWhileShorteningServerResponse) {
			logrus.Errorf("Error while shortening response from the target server: %s", err.Error())
		} else {
			logrus.Errorf("Unknown error while shortening response: %s", err.Error())
		}
		return nil, err
	}
	return result, nil
}

func (processor *ResponseShorteningService) getRules(api *shortreq.ShortenedAPI) map[string]string {
	resultRules := make(map[string]string)
	for _, rule := range api.ShorteningRules {
		resultRules[rule.FieldName] = rule.FieldValueQuery
	}
	return resultRules
}

func NewResponseShorteningService(
	configDAO shortreq.IOutgoingRequestConfigDAO,
	headersDAO shortreq.IOutgoingRequestHeaderDAO,
	paramsDAO shortreq.IOutgoingRequestParamDAO,
	shortener shortener.IResponseShortener,
	limiter ILoopLimiter,
) *ResponseShorteningService {
	return &ResponseShorteningService{
		configDAO:  configDAO,
		headersDAO: headersDAO,
		paramsDAO:  paramsDAO,
		shortener:  shortener,
		limiter:    limiter,
	}
}
