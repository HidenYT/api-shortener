package security

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	API_AUTH_TOKEN_QUERY_KEY = "token"
	API_AUTH_TOKEN_ENV_KEY   = "API_KEY"
)

func APITokenChecker() gin.HandlerFunc {
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
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}
}
