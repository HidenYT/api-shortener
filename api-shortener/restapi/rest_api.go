package restapi

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AttachRESTAPIGroup(r *gin.Engine, RESTService IRESTService) {
	restGroup := r.Group("/rest")

	apiGroup := restGroup.Group("/api")
	apiGroup.POST("", func(c *gin.Context) {
		createShortenedAPI(c, RESTService)
	})
	apiGroup.DELETE("/:id", func(c *gin.Context) {
		deleteShortenedAPI(c, RESTService)
	})

	configGroup := restGroup.Group("/configs")
	configGroup.POST("", func(c *gin.Context) {
		createOutgoingRequestConfig(c, RESTService)
	})
	configGroup.GET("/:id", func(c *gin.Context) {
		getOutgoingRequestConfig(c, RESTService)
	})
	configGroup.GET("/", func(c *gin.Context) {
		getOutgoingRequestConfigByAPIID(c, RESTService)
	})
	configGroup.PUT("/:id", func(c *gin.Context) {
		updateOutgoingRequestConfig(c, RESTService)
	})
	configGroup.DELETE("/:id", func(c *gin.Context) {
		deleteOutgoingRequestConfig(c, RESTService)
	})

	rulesGroup := restGroup.Group("/rules")
	rulesGroup.POST("", func(c *gin.Context) {
		createShorteningRule(c, RESTService)
	})
	rulesGroup.GET("/:id", func(c *gin.Context) {
		getShorteningRule(c, RESTService)
	})
	rulesGroup.GET("/", func(c *gin.Context) {
		getAllShorteningRulesByAPIID(c, RESTService)
	})
	rulesGroup.PUT("/:id", func(c *gin.Context) {
		updateShorteningRule(c, RESTService)
	})
	rulesGroup.DELETE("/:id", func(c *gin.Context) {
		deleteShorteningRule(c, RESTService)
	})

	headersGroup := restGroup.Group("/headers")
	headersGroup.POST("", func(c *gin.Context) {
		createOutgoingRequestHeader(c, RESTService)
	})
	headersGroup.GET("/:id", func(c *gin.Context) {
		getOutgoingRequestHeader(c, RESTService)
	})
	headersGroup.GET("/", func(c *gin.Context) {
		getAllOutgoingRequestHeadersByConfigID(c, RESTService)
	})
	headersGroup.PUT("/:id", func(c *gin.Context) {
		updateOutgoingRequestHeader(c, RESTService)
	})
	headersGroup.DELETE("/:id", func(c *gin.Context) {
		deleteOutgoingRequestHeader(c, RESTService)
	})

	paramsGroup := restGroup.Group("/params")
	paramsGroup.POST("", func(c *gin.Context) {
		createOutgoingRequestParam(c, RESTService)
	})
	paramsGroup.GET("/:id", func(c *gin.Context) {
		getOutgoingRequestParam(c, RESTService)
	})
	paramsGroup.GET("/", func(c *gin.Context) {
		getAllOutgoingRequestParamsByConfigID(c, RESTService)
	})
	paramsGroup.PUT("/:id", func(c *gin.Context) {
		updateOutgoingRequestParam(c, RESTService)
	})
	paramsGroup.DELETE("/:id", func(c *gin.Context) {
		deleteOutgoingRequestParam(c, RESTService)
	})
}

func getUintFromPath(name string, c *gin.Context) (uint, error) {
	uintRaw, err := strconv.ParseUint(c.Param(name), 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(uintRaw), nil
}

func getUintFromQuery(name string, c *gin.Context) (uint, error) {
	apiIdRaw, err := strconv.ParseUint(c.Query(name), 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(apiIdRaw), nil
}

// API

func createShortenedAPI(c *gin.Context, restService IRESTService) {
	res, err := restService.CreateAPI()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func deleteShortenedAPI(c *gin.Context, restService IRESTService) {
	apiId, err := getUintFromPath("id", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = restService.DeleteAPI(apiId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// RequestConfig

func createOutgoingRequestConfig(c *gin.Context, restService IRESTService) {
	var config OutgoingRequestConfigRequest
	if err := c.BindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := restService.CreateConfig(&config)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func getOutgoingRequestConfig(c *gin.Context, restService IRESTService) {
	configId, err := getUintFromPath("id", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := restService.GetConfig(configId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func getOutgoingRequestConfigByAPIID(c *gin.Context, restService IRESTService) {
	apiId, err := getUintFromQuery("apiID", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := restService.GetConfigByAPIID(apiId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func updateOutgoingRequestConfig(c *gin.Context, restService IRESTService) {
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
	res, err := restService.UpdateConfig(configId, &config)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func deleteOutgoingRequestConfig(c *gin.Context, restService IRESTService) {
	configId, err := getUintFromPath("id", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = restService.DeleteConfig(configId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// ShorteningRule

func createShorteningRule(c *gin.Context, restService IRESTService) {
	var rule ShorteningRuleRequest
	err := c.BindJSON(&rule)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := restService.CreateRule(&rule)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func getShorteningRule(c *gin.Context, restService IRESTService) {
	ruleId, err := getUintFromPath("id", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := restService.GetRule(ruleId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func getAllShorteningRulesByAPIID(c *gin.Context, restService IRESTService) {
	apiId, err := getUintFromQuery("apiID", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := restService.GetAllRulesByAPIID(apiId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func updateShorteningRule(c *gin.Context, restService IRESTService) {
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
	res, err := restService.UpdateRule(ruleId, &rule)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func deleteShorteningRule(c *gin.Context, restService IRESTService) {
	ruleId, err := getUintFromPath("id", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = restService.DeleteRule(ruleId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// RequestHeader

func createOutgoingRequestHeader(c *gin.Context, restService IRESTService) {
	var header OutgoingRequestHeaderRequest
	err := c.BindJSON(&header)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := restService.CreateRequestHeader(&header)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func getOutgoingRequestHeader(c *gin.Context, restService IRESTService) {
	headerId, err := getUintFromPath("id", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := restService.GetRequestHeader(headerId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func getAllOutgoingRequestHeadersByConfigID(c *gin.Context, restService IRESTService) {
	configId, err := getUintFromQuery("configID", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := restService.GetAllRequestHeadersByConfigID(configId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func updateOutgoingRequestHeader(c *gin.Context, restService IRESTService) {
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
	res, err := restService.UpdateRequestHeader(headerId, &header)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func deleteOutgoingRequestHeader(c *gin.Context, restService IRESTService) {
	headerId, err := getUintFromPath("id", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = restService.DeleteRequestHeader(headerId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// RequestParam

func createOutgoingRequestParam(c *gin.Context, restService IRESTService) {
	var param OutgoingRequestParamRequest
	err := c.BindJSON(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := restService.CreateRequestParam(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func getOutgoingRequestParam(c *gin.Context, restService IRESTService) {
	paramId, err := getUintFromPath("id", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := restService.GetRequestParam(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func getAllOutgoingRequestParamsByConfigID(c *gin.Context, restService IRESTService) {
	configId, err := getUintFromQuery("configID", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := restService.GetAllRequestParamsByConfigID(configId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func updateOutgoingRequestParam(c *gin.Context, restService IRESTService) {
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
	res, err := restService.UpdateRequestParam(paramId, &param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func deleteOutgoingRequestParam(c *gin.Context, restService IRESTService) {
	paramId, err := getUintFromPath("id", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = restService.DeleteRequestParam(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
