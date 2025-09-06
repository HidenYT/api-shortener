package shortener

import "net/http"

type ShortenedResponse struct {
	JSON       *map[string]any
	StatusCode int
	Headers    http.Header
}
