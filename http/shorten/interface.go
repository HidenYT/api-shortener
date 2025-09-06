package http_shortener

import (
	shortener "github.com/HidenYT/api-shortener/response-shortener"
	"github.com/HidenYT/api-shortener/shortreq"
)

type IResponseShorteningService interface {
	ProcessRequest(api *shortreq.ShortenedAPI) (*shortener.ShortenedResponse, error)
}
