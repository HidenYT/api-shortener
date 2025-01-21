package shortreq

import (
	"github.com/gin-gonic/gin"
)

func shorteningView(c *gin.Context, shorteningService IResponseShorteningService) {
	api := c.MustGet(CTX_API_KEY)
	shorteningService.ProcessRequest(api.(*ShortenedAPI), c)
}

func AttachAPIShorteningGroup(r *gin.Engine, shorteningService IResponseShorteningService, apiRepo IShortenedAPIDAO) {
	apiGroup := r.Group("/api")

	apiGroup.Use(APIAuthChecker(apiRepo))
	apiGroup.Any("/:apiID", func(c *gin.Context) {
		shorteningView(c, shorteningService)
	})
}
