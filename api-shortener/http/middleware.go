package http

import (
	"api-shortener/shortreq"
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	CTX_API_KEY              = "api"
	API_AUTH_TOKEN_QUERY_KEY = "token"
	API_AUTH_TOKEN_ENV_KEY   = "API_KEY"
)

var (
	errUnathorized                = errors.New("Unauthorized")
	errAPIIDNotFoundInRequestPath = errors.New("API ID not found in request path")
	errInvalidAPIID               = errors.New("API ID is invalid")
)

func apiTokenChecker() gin.HandlerFunc {
	return func(c *gin.Context) {
		passedToken := c.Query(API_AUTH_TOKEN_QUERY_KEY)
		realToken, ok := os.LookupEnv(API_AUTH_TOKEN_ENV_KEY)
		if !ok {
			panic("No API token found in envs")
		}
		if passedToken == realToken {
			c.Next()
			return
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": errUnathorized.Error()})
	}
}

func apiIDChecker(apiDAO shortreq.IShortenedAPIDAO) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiID, err := strconv.ParseUint(c.Param("apiID"), 10, 32)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errAPIIDNotFoundInRequestPath.Error()})
			return
		}

		api, err := apiDAO.Get(uint(apiID))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errInvalidAPIID.Error()})
			return
		}
		c.Set(CTX_API_KEY, api)
		c.Next()
	}
}
