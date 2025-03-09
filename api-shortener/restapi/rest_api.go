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
		getAllOutgoingRequestConfigsByAPIID(c, RESTService)
	})
	configGroup.PUT("", func(c *gin.Context) {
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
	rulesGroup.PUT("", func(c *gin.Context) {
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
	headersGroup.PUT("", func(c *gin.Context) {
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
	paramsGroup.PUT("", func(c *gin.Context) {
		updateOutgoingRequestParam(c, RESTService)
	})
	paramsGroup.DELETE("/:id", func(c *gin.Context) {
		deleteOutgoingRequestParam(c, RESTService)
	})
}

// API

func createShortenedAPI(c *gin.Context, restService IRESTService) {
	var shortenedAPI ShortenedAPI
	err := c.BindJSON(&shortenedAPI)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := restService.CreateAPI(&shortenedAPI)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func deleteShortenedAPI(c *gin.Context, restService IRESTService) {
	apiIdRaw, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	apiId := uint(apiIdRaw)
	err = restService.DeleteAPI(apiId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// RequestConfig

func createOutgoingRequestConfig(c *gin.Context, restService IRESTService) {
	var config OutgoingRequestConfig
	err := c.BindJSON(&config)
	if err != nil {
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
	configIdRaw, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	configId := uint(configIdRaw)
	res, err := restService.GetConfig(configId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func getAllOutgoingRequestConfigsByAPIID(c *gin.Context, restService IRESTService) {
	apiIdRaw, err := strconv.ParseUint(c.Query("apiID"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	apiId := uint(apiIdRaw)
	res, err := restService.GetAllConfigsByAPIID(apiId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func updateOutgoingRequestConfig(c *gin.Context, restService IRESTService) {
	var config OutgoingRequestConfig
	err := c.BindJSON(&config)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := restService.UpdateConfig(&config)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func deleteOutgoingRequestConfig(c *gin.Context, restService IRESTService) {
	configIdRaw, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	configId := uint(configIdRaw)
	err = restService.DeleteConfig(configId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// ShorteningRule

func createShorteningRule(c *gin.Context, restService IRESTService) {
	var rule ShorteningRule
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
	ruleIdRaw, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ruleId := uint(ruleIdRaw)
	res, err := restService.GetRule(ruleId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func getAllShorteningRulesByAPIID(c *gin.Context, restService IRESTService) {
	apiIdRaw, err := strconv.ParseUint(c.Query("apiID"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	apiId := uint(apiIdRaw)
	res, err := restService.GetAllRulesByAPIID(apiId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func updateShorteningRule(c *gin.Context, restService IRESTService) {
	var rule ShorteningRule
	err := c.BindJSON(&rule)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := restService.UpdateRule(&rule)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func deleteShorteningRule(c *gin.Context, restService IRESTService) {
	ruleIdRaw, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ruleId := uint(ruleIdRaw)
	err = restService.DeleteRule(ruleId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// RequestHeader

func createOutgoingRequestHeader(c *gin.Context, restService IRESTService) {
	var header OutgoingRequestHeader
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
	headerIdRaw, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	headerId := uint(headerIdRaw)
	res, err := restService.GetRequestHeader(headerId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func getAllOutgoingRequestHeadersByConfigID(c *gin.Context, restService IRESTService) {
	configIdRaw, err := strconv.ParseUint(c.Query("configID"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	configId := uint(configIdRaw)
	res, err := restService.GetAllRequestHeadersByConfigID(configId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func updateOutgoingRequestHeader(c *gin.Context, restService IRESTService) {
	var header OutgoingRequestHeader
	err := c.BindJSON(&header)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := restService.UpdateRequestHeader(&header)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func deleteOutgoingRequestHeader(c *gin.Context, restService IRESTService) {
	headerIdRaw, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	headerId := uint(headerIdRaw)
	err = restService.DeleteRequestHeader(headerId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// RequestParam

func createOutgoingRequestParam(c *gin.Context, restService IRESTService) {
	var param OutgoingRequestParam
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
	paramIdRaw, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	paramId := uint(paramIdRaw)
	res, err := restService.GetRequestParam(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func getAllOutgoingRequestParamsByConfigID(c *gin.Context, restService IRESTService) {
	configIdRaw, err := strconv.ParseUint(c.Query("configID"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	configId := uint(configIdRaw)
	res, err := restService.GetAllRequestParamsByConfigID(configId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func updateOutgoingRequestParam(c *gin.Context, restService IRESTService) {
	var param OutgoingRequestParam
	err := c.BindJSON(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := restService.UpdateRequestParam(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func deleteOutgoingRequestParam(c *gin.Context, restService IRESTService) {
	paramIdRaw, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	paramId := uint(paramIdRaw)
	err = restService.DeleteRequestParam(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
