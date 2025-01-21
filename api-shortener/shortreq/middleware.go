package shortreq

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	CTX_API_KEY   = "api"
	CTX_ERROR_KEY = "error"
)

type APINotFoundError struct {
	apiID uint
}

func (e *APINotFoundError) Error() string {
	return fmt.Sprintf("API with id %d not found", e.apiID)
}

func APIAuthChecker(apiRepo IShortenedAPIDAO) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiID, err := strconv.ParseUint(c.Param("apiID"), 10, 32)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		api, err := apiRepo.Get(uint(apiID))
		if err != nil {
			err = &APINotFoundError{apiID: uint(apiID)}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.Set(CTX_API_KEY, api)
		c.Next()
	}
}
