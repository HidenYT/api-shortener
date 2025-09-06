package http_shortener

import (
	shortener "github.com/HidenYT/api-shortener/internal/response-shortener"
	db_model "github.com/HidenYT/api-shortener/internal/storage/db-model/api"
)

type IResponseShorteningService interface {
	ProcessRequest(api *db_model.ShortenedAPI) (*shortener.ShortenedResponse, error)
}
