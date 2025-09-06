package http_shortener

import (
	"errors"
	"net/http"

	http_common "github.com/HidenYT/api-shortener/internal/http/common"
	shortener "github.com/HidenYT/api-shortener/internal/response-shortener"
	api_dao "github.com/HidenYT/api-shortener/internal/storage/dao/api"
	db_model "github.com/HidenYT/api-shortener/internal/storage/db-model/api"
	"github.com/gin-gonic/gin"
)

func shorteningView(c *gin.Context, shorteningService IResponseShorteningService) {
	api := c.MustGet(http_common.CTX_API_KEY)
	response, err := shorteningService.ProcessRequest(api.(*db_model.ShortenedAPI))
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

func AttachAPIShorteningGroup(r *gin.Engine, shorteningService IResponseShorteningService, apiDAO api_dao.IShortenedAPIDAO) {
	apiGroup := r.Group("/api")

	apiGroup.Use(http_common.APIIDChecker(apiDAO))
	apiGroup.Any("/:apiID", func(c *gin.Context) {
		shorteningView(c, shorteningService)
	})
}
