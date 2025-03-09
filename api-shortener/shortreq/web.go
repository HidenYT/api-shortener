package shortreq

import (
	"api-shortener/restapi"

	"github.com/gin-gonic/gin"
)

func shorteningView(c *gin.Context, shorteningService IResponseShorteningService) {
	api := c.MustGet(CTX_API_KEY)
	shorteningService.ProcessRequest(api.(*restapi.ShortenedAPI), c)
}

func AttachAPIShorteningGroup(r *gin.Engine, shorteningService IResponseShorteningService, apiRepo restapi.IShortenedAPIDAO) {
	apiGroup := r.Group("/api")

	apiGroup.Use(APIAuthChecker(apiRepo))
	apiGroup.Any("/:apiID", func(c *gin.Context) {
		shorteningView(c, shorteningService)
	})
}
