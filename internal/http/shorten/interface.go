package http_shortener

import (
	shortener "github.com/HidenYT/api-shortener/internal/response-shortener"
	api_dao "github.com/HidenYT/api-shortener/internal/storage/dao"
)

type IResponseShorteningService interface {
	ProcessRequest(api *api_dao.ShortenedAPI) (*shortener.ShortenedResponse, error)
}
