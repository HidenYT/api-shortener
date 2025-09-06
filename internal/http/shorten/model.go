package http_shortener

import shortener "github.com/HidenYT/api-shortener/internal/response-shortener"

type ShortenedResponseMeta struct {
	Err string `json:"error"`
}

type ShortenedAPIResponse struct {
	Meta *ShortenedResponseMeta `json:"meta,omitempty"`
	Data *map[string]any        `json:"data,omitempty"`
}

func shortenedAPIResponseFromError(err error) ShortenedAPIResponse {
	return ShortenedAPIResponse{
		Meta: &ShortenedResponseMeta{
			Err: err.Error(),
		},
	}
}

func shortenedAPIResponseFromResponse(response *shortener.ShortenedResponse) ShortenedAPIResponse {
	return ShortenedAPIResponse{
		Data: response.JSON,
	}
}
