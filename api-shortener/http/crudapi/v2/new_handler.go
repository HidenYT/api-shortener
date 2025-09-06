package http

import (
	http_common "api-shortener/http/common"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func AttachHandlerGroupV2(
	r *gin.Engine,
	service IAPIDTOService,
) {
	apiGroup := r.Group("/rest/api/v2")
	apiGroup.POST("", func(c *gin.Context) {
		createAPIHandle(c, service)
	})
	apiGroup.GET("/:id", func(c *gin.Context) {
		getAPIHandle(c, service)
	})
	apiGroup.PUT("/:id", func(c *gin.Context) {
		updateAPIHandle(c, service)
	})
	apiGroup.DELETE("/:id", func(c *gin.Context) {
		deleteAPIHandle(c, service)
	})
}

func createAPIHandle(c *gin.Context, service IAPIDTOService) {
	var dto ShortenedAPIDTO
	if err := c.BindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := service.Create(&dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func getAPIHandle(c *gin.Context, service IAPIDTOService) {
	apiID, err := http_common.GetUintFromPath("id", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := service.GetByID(apiID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func updateAPIHandle(c *gin.Context, service IAPIDTOService) {
	apiID, err := http_common.GetUintFromPath("id", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var dto ShortenedAPIDTO
	if err := c.BindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := service.Update(apiID, &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func deleteAPIHandle(c *gin.Context, service IAPIDTOService) {
	apiID, err := http_common.GetUintFromPath("id", c)
	logrus.Info("Delete", apiID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = service.DeleteByID(apiID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
