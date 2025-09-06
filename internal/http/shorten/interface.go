package http_shortener

import (
	shortener "github.com/HidenYT/api-shortener/internal/response-shortener"
	"github.com/HidenYT/api-shortener/internal/shortreq"
)

type IResponseShorteningService interface {
	ProcessRequest(api *shortreq.ShortenedAPI) (*shortener.ShortenedResponse, error)
}
