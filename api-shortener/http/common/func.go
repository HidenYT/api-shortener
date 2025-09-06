package http

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUintFromPath(name string, c *gin.Context) (uint, error) {
	uintRaw, err := strconv.ParseUint(c.Param(name), 10, 32)
	if err != nil {
		return 0, fmt.Errorf("unable to parse %s from path", name)
	}
	return uint(uintRaw), nil
}

func GetUintFromQuery(name string, c *gin.Context) (uint, error) {
	apiIdRaw, err := strconv.ParseUint(c.Query(name), 10, 32)
	if err != nil {
		return 0, fmt.Errorf("unable to parse %s from query", name)
	}
	return uint(apiIdRaw), nil
}
