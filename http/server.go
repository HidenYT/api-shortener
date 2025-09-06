package http

import (
	http_common "github.com/HidenYT/api-shortener/http/common"
	crudapi_v1 "github.com/HidenYT/api-shortener/http/crudapi/v1"
	crudapi_v2 "github.com/HidenYT/api-shortener/http/crudapi/v2"
	"github.com/HidenYT/api-shortener/shortreq"

	"github.com/gin-gonic/gin"
)

func NewHTTPServer(
	apiDAO shortreq.IShortenedAPIDAO,
	shorteningService crudapi_v1.IResponseShorteningService,
	apiService crudapi_v1.IAPIService,
	configService crudapi_v1.IRequestConfigService,
	headerService crudapi_v1.IRequestHeaderService,
	paramService crudapi_v1.IRequestParamService,
	rulesService crudapi_v1.IShorteningRuleService,
	apiDTOService crudapi_v2.IAPIDTOService,
) *gin.Engine {
	ginServer := gin.Default()
	ginServer.Use(http_common.APITokenChecker())
	crudapi_v1.AttachAPIShorteningGroup(ginServer, shorteningService, apiDAO)
	crudapi_v1.AttachRESTAPIGroup(ginServer, apiService, configService, headerService, paramService, rulesService)
	crudapi_v2.AttachHandlerGroupV2(ginServer, apiDTOService)
	return ginServer
}
