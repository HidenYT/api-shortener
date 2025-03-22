package shortener

import (
	"net/http"
)

type IResponseShortener interface {
	ProcessRequest(request *http.Request, rules map[string]string) (*ShortenedResponse, error)
}

type IOutgoingRequestClient interface {
	MakeRequest(request *http.Request) (*http.Response, error)
}
