package restapi

import (
	"api-shortener/shortreq"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AttachRESTAPIGroup(r *gin.Engine, RESTService IRESTService) {
	apiGroup := r.Group("/rest")

	apiGroup.POST("", func(c *gin.Context) {
		createShortenedAPI(c, RESTService)
	})
	apiGroup.GET("/:id", func(c *gin.Context) {
		getShortenedAPI(c, RESTService)
	})
	apiGroup.PUT("", func(c *gin.Context) {
		updateShortenedAPI(c, RESTService)
	})
	apiGroup.DELETE("/:id", func(c *gin.Context) {
		deleteShortenedAPI(c, RESTService)
	})
}

func createShortenedAPI(c *gin.Context, restService IRESTService) {
	var shortenedAPI shortreq.ShortenedAPI
	err := c.BindJSON(&shortenedAPI)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := restService.Create(&shortenedAPI)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func getShortenedAPI(c *gin.Context, restService IRESTService) {
	apiIdRaw, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	apiId := uint(apiIdRaw)
	res, err := restService.Get(apiId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func updateShortenedAPI(c *gin.Context, restService IRESTService) {
	var shortenedAPI shortreq.ShortenedAPI
	err := c.BindJSON(&shortenedAPI)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := restService.Update(&shortenedAPI)
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
	err = restService.Delete(apiId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
