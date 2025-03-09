package shortreq

import (
	"api-shortener/restapi"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type IResponseShorteningService interface {
	ProcessRequest(api *restapi.ShortenedAPI, c *gin.Context)
}

type ResponseShorteningService struct {
	incomingRequestProcessor IIncomingRequestProcessor
	outgoingRequestProcessor IOutgoingRequestProcessor
	limiter                  ILoopLimiter
}

type RequestAlreadySentError struct {
	apiId uint
}

func (e *RequestAlreadySentError) Error() string {
	return fmt.Sprintf("Request is already sent to the API %d", e.apiId)
}

func (s *ResponseShorteningService) ProcessRequest(api *restapi.ShortenedAPI, c *gin.Context) {
	if !s.limiter.AddNewRequest(api.ID) {
		err := &RequestAlreadySentError{apiId: api.ID}
		c.JSON(http.StatusTooManyRequests, gin.H{"error": err.Error()})
		logrus.Warningf("Max requests limit exceeded for API %d", api.ID)
		return
	}
	defer s.limiter.RemoveRequest(api.ID)

	request, err := s.incomingRequestProcessor.CreateOutgoingRequest(api)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logrus.Errorf("Error while creating request for API %d: %s", api.ID, err.Error())
		return
	}
	s.outgoingRequestProcessor.Process(request, c, api)
}

func NewResponseShorteningService(
	incomingRequestProcessor IIncomingRequestProcessor,
	outgoingRequestProcessor IOutgoingRequestProcessor,
	limiter ILoopLimiter,
) IResponseShorteningService {
	return &ResponseShorteningService{
		incomingRequestProcessor: incomingRequestProcessor,
		outgoingRequestProcessor: outgoingRequestProcessor,
		limiter:                  limiter,
	}
}
