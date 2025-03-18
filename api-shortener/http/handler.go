package http

import (
	"api-shortener/shortreq"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func attachRESTAPIGroup(
	r *gin.Engine,
	apiService IAPIService,
	configService IRequestConfigService,
	headerService IRequestHeaderService,
	paramService IRequestParamService,
	rulesService IShorteningRuleService,
) {
	restGroup := r.Group("/rest")

	apiGroup := restGroup.Group("/api")
	apiGroup.POST("", func(c *gin.Context) {
		createShortenedAPI(c, apiService)
	})
	apiGroup.DELETE("/:id", func(c *gin.Context) {
		deleteShortenedAPI(c, apiService)
	})

	configGroup := restGroup.Group("/configs")
	configGroup.POST("", func(c *gin.Context) {
		createOutgoingRequestConfig(c, configService)
	})
	configGroup.GET("/:id", func(c *gin.Context) {
		getOutgoingRequestConfig(c, configService)
	})
	configGroup.GET("/", func(c *gin.Context) {
		getOutgoingRequestConfigByAPIID(c, configService)
	})
	configGroup.PUT("/:id", func(c *gin.Context) {
		updateOutgoingRequestConfig(c, configService)
	})
	configGroup.DELETE("/:id", func(c *gin.Context) {
		deleteOutgoingRequestConfig(c, configService)
	})

	rulesGroup := restGroup.Group("/rules")
	rulesGroup.POST("", func(c *gin.Context) {
		createShorteningRule(c, rulesService)
	})
	rulesGroup.GET("/:id", func(c *gin.Context) {
		getShorteningRule(c, rulesService)
	})
	rulesGroup.GET("/", func(c *gin.Context) {
		getAllShorteningRulesByAPIID(c, rulesService)
	})
	rulesGroup.PUT("/:id", func(c *gin.Context) {
		updateShorteningRule(c, rulesService)
	})
	rulesGroup.DELETE("/:id", func(c *gin.Context) {
		deleteShorteningRule(c, rulesService)
	})

	headersGroup := restGroup.Group("/headers")
	headersGroup.POST("", func(c *gin.Context) {
		createOutgoingRequestHeader(c, headerService)
	})
	headersGroup.GET("/:id", func(c *gin.Context) {
		getOutgoingRequestHeader(c, headerService)
	})
	headersGroup.GET("/", func(c *gin.Context) {
		getAllOutgoingRequestHeadersByConfigID(c, headerService)
	})
	headersGroup.PUT("/:id", func(c *gin.Context) {
		updateOutgoingRequestHeader(c, headerService)
	})
	headersGroup.DELETE("/:id", func(c *gin.Context) {
		deleteOutgoingRequestHeader(c, headerService)
	})

	paramsGroup := restGroup.Group("/params")
	paramsGroup.POST("", func(c *gin.Context) {
		createOutgoingRequestParam(c, paramService)
	})
	paramsGroup.GET("/:id", func(c *gin.Context) {
		getOutgoingRequestParam(c, paramService)
	})
	paramsGroup.GET("/", func(c *gin.Context) {
		getAllOutgoingRequestParamsByConfigID(c, paramService)
	})
	paramsGroup.PUT("/:id", func(c *gin.Context) {
		updateOutgoingRequestParam(c, paramService)
	})
	paramsGroup.DELETE("/:id", func(c *gin.Context) {
		deleteOutgoingRequestParam(c, paramService)
	})
}

func getUintFromPath(name string, c *gin.Context) (uint, error) {
	uintRaw, err := strconv.ParseUint(c.Param(name), 10, 32)
	if err != nil {
		return 0, fmt.Errorf("unable to parse %s from path", name)
	}
	return uint(uintRaw), nil
}

func getUintFromQuery(name string, c *gin.Context) (uint, error) {
	apiIdRaw, err := strconv.ParseUint(c.Query(name), 10, 32)
	if err != nil {
		return 0, fmt.Errorf("unable to parse %s from query", name)
	}
	return uint(apiIdRaw), nil
}

// API

func createShortenedAPI(c *gin.Context, apiService IAPIService) {
	res, err := apiService.Create()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func deleteShortenedAPI(c *gin.Context, apiService IAPIService) {
	apiId, err := getUintFromPath("id", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = apiService.Delete(apiId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// RequestConfig

func createOutgoingRequestConfig(c *gin.Context, configService IRequestConfigService) {
	var config OutgoingRequestConfigRequest
	if err := c.BindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := configService.Create(&config)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func getOutgoingRequestConfig(c *gin.Context, configService IRequestConfigService) {
	configId, err := getUintFromPath("id", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := configService.GetByID(configId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func getOutgoingRequestConfigByAPIID(c *gin.Context, configService IRequestConfigService) {
	apiId, err := getUintFromQuery("apiID", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := configService.GetByAPIID(apiId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func updateOutgoingRequestConfig(c *gin.Context, configService IRequestConfigService) {
	configId, err := getUintFromPath("id", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var config OutgoingRequestConfigRequest
	if err = c.BindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := configService.Update(configId, &config)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func deleteOutgoingRequestConfig(c *gin.Context, configService IRequestConfigService) {
	configId, err := getUintFromPath("id", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = configService.Delete(configId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// ShorteningRule

func createShorteningRule(c *gin.Context, rulesService IShorteningRuleService) {
	var rule ShorteningRuleRequest
	err := c.BindJSON(&rule)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := rulesService.Create(&rule)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func getShorteningRule(c *gin.Context, rulesService IShorteningRuleService) {
	ruleId, err := getUintFromPath("id", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := rulesService.GetByID(ruleId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func getAllShorteningRulesByAPIID(c *gin.Context, rulesService IShorteningRuleService) {
	apiId, err := getUintFromQuery("apiID", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := rulesService.GetAllByAPIID(apiId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func updateShorteningRule(c *gin.Context, rulesService IShorteningRuleService) {
	ruleId, err := getUintFromPath("id", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var rule ShorteningRuleRequest
	if err = c.BindJSON(&rule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := rulesService.Update(ruleId, &rule)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func deleteShorteningRule(c *gin.Context, rulesService IShorteningRuleService) {
	ruleId, err := getUintFromPath("id", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = rulesService.Delete(ruleId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// RequestHeader

func createOutgoingRequestHeader(c *gin.Context, headerService IRequestHeaderService) {
	var header OutgoingRequestHeaderRequest
	err := c.BindJSON(&header)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := headerService.Create(&header)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func getOutgoingRequestHeader(c *gin.Context, headerService IRequestHeaderService) {
	headerId, err := getUintFromPath("id", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := headerService.GetByID(headerId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func getAllOutgoingRequestHeadersByConfigID(c *gin.Context, headerService IRequestHeaderService) {
	configId, err := getUintFromQuery("configID", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := headerService.GetAllByConfigID(configId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func updateOutgoingRequestHeader(c *gin.Context, headerService IRequestHeaderService) {
	headerId, err := getUintFromPath("id", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var header OutgoingRequestHeaderRequest
	if err = c.BindJSON(&header); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := headerService.Update(headerId, &header)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func deleteOutgoingRequestHeader(c *gin.Context, headerService IRequestHeaderService) {
	headerId, err := getUintFromPath("id", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = headerService.Delete(headerId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// RequestParam

func createOutgoingRequestParam(c *gin.Context, paramService IRequestParamService) {
	var param OutgoingRequestParamRequest
	err := c.BindJSON(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := paramService.Create(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func getOutgoingRequestParam(c *gin.Context, paramService IRequestParamService) {
	paramId, err := getUintFromPath("id", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := paramService.GetByID(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func getAllOutgoingRequestParamsByConfigID(c *gin.Context, paramService IRequestParamService) {
	configId, err := getUintFromQuery("configID", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := paramService.GetAllByConfigID(configId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func updateOutgoingRequestParam(c *gin.Context, paramService IRequestParamService) {
	paramId, err := getUintFromPath("id", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var param OutgoingRequestParamRequest
	if err = c.BindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := paramService.Update(paramId, &param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func deleteOutgoingRequestParam(c *gin.Context, paramService IRequestParamService) {
	paramId, err := getUintFromPath("id", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = paramService.Delete(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func shorteningView(c *gin.Context, shorteningService IResponseShorteningService) {
	api := c.MustGet(CTX_API_KEY)
	response, err := shorteningService.ProcessRequest(api.(*shortreq.ShortenedAPI))
	if err != nil {
		if errors.Is(err, errRequestIsAlreadySent) {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": err.Error()})
		} else if errors.Is(err, errWhileShorteningServerResponse) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	for header := range response.headers {
		c.Writer.Header().Add(header, response.headers.Get(header))
	}
	c.JSON(response.statusCode, response.json)
}

func attachAPIShorteningGroup(r *gin.Engine, shorteningService IResponseShorteningService, apiDAO shortreq.IShortenedAPIDAO) {
	apiGroup := r.Group("/api")

	apiGroup.Use(apiIDChecker(apiDAO))
	apiGroup.Any("/:apiID", func(c *gin.Context) {
		shorteningView(c, shorteningService)
	})
}
