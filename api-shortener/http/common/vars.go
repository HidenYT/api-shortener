package http

import "errors"

const (
	CTX_API_KEY              = "api"
	API_AUTH_TOKEN_QUERY_KEY = "token"
	API_AUTH_TOKEN_ENV_KEY   = "API_KEY"
)

var (
	errUnathorized                = errors.New("Unauthorized")
	errAPIIDNotFoundInRequestPath = errors.New("API ID not found in request path")
	errAPIIDNotFound              = errors.New("API ID not found")
)
