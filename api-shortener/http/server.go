package http

import (
	"api-shortener/shortreq"

	"github.com/gin-gonic/gin"
)

func NewHTTPServer(
	apiDAO shortreq.IShortenedAPIDAO,
	shorteningService IResponseShorteningService,
	apiService IAPIService,
	configService IRequestConfigService,
	headerService IRequestHeaderService,
	paramService IRequestParamService,
	rulesService IShorteningRuleService,
) *gin.Engine {
	ginServer := gin.Default()
	ginServer.Use(apiTokenChecker())
	attachAPIShorteningGroup(ginServer, shorteningService, apiDAO)
	attachRESTAPIGroup(ginServer, apiService, configService, headerService, paramService, rulesService)
	return ginServer
}
