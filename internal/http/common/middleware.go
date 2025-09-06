package http

import (
	"net/http"
	"os"
	"strconv"

	"github.com/HidenYT/api-shortener/internal/shortreq"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func APITokenChecker() gin.HandlerFunc {
	return func(c *gin.Context) {
		passedToken := c.Query(API_AUTH_TOKEN_QUERY_KEY)
		realToken, ok := os.LookupEnv(API_AUTH_TOKEN_ENV_KEY)
		if !ok {
			logrus.Fatalf("API_AUTH_TOKEN_ENV_KEY not found in env")
		}
		if passedToken == realToken {
			c.Next()
			return
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": errUnathorized.Error()})
	}
}

func APIIDChecker(apiDAO shortreq.IShortenedAPIDAO) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiID, err := strconv.ParseUint(c.Param("apiID"), 10, 32)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errAPIIDNotFoundInRequestPath.Error()})
			return
		}

		api, err := apiDAO.Get(uint(apiID))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errAPIIDNotFound.Error()})
			return
		}
		c.Set(CTX_API_KEY, api)
		c.Next()
	}
}
