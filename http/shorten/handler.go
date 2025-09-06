package http_shortener

import (
	"errors"
	"net/http"

	http_common "github.com/HidenYT/api-shortener/http/common"
	shortener "github.com/HidenYT/api-shortener/response-shortener"
	"github.com/HidenYT/api-shortener/shortreq"
	"github.com/gin-gonic/gin"
)

func shorteningView(c *gin.Context, shorteningService IResponseShorteningService) {
	api := c.MustGet(http_common.CTX_API_KEY)
	response, err := shorteningService.ProcessRequest(api.(*shortreq.ShortenedAPI))
	if err != nil {
		var status int
		if errors.Is(err, errRequestIsAlreadySent) {
			status = http.StatusTooManyRequests
		} else if errors.Is(err, shortener.ErrWhileShorteningServerResponse) {
			status = http.StatusBadRequest
		} else {
			status = http.StatusInternalServerError
		}
		c.JSON(status, shortenedAPIResponseFromError(err))
		return
	}
	for header := range response.Headers {
		c.Writer.Header().Add(header, response.Headers.Get(header))
	}
	c.JSON(response.StatusCode, shortenedAPIResponseFromResponse(response))
}

func AttachAPIShorteningGroup(r *gin.Engine, shorteningService IResponseShorteningService, apiDAO shortreq.IShortenedAPIDAO) {
	apiGroup := r.Group("/api")

	apiGroup.Use(http_common.APIIDChecker(apiDAO))
	apiGroup.Any("/:apiID", func(c *gin.Context) {
		shorteningView(c, shorteningService)
	})
}
