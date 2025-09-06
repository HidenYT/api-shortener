package shortener

import (
	"errors"
	"io"
	"net/http"

	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/oj"
)

var (
	ErrWhileMakingRequest            = errors.New("error while making request to the server")
	ErrWhileReadingServerResponse    = errors.New("error while reading server response")
	ErrWhileShorteningServerResponse = errors.New("error while shortening server response")
)

type ResponseShortener struct {
	client IOutgoingRequestClient
}

func NewResponseShortener(client IOutgoingRequestClient) *ResponseShortener {
	return &ResponseShortener{client: client}
}

func (s *ResponseShortener) ProcessRequest(request *http.Request, rules map[string]string) (*ShortenedResponse, error) {
	response, err := s.client.MakeRequest(request)
	if err != nil {
		return nil, ErrWhileMakingRequest
	}

	body, err := io.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, ErrWhileReadingServerResponse
	}

	resultHeader := make(http.Header)
	for headerName := range response.Header {
		resultHeader.Add(headerName, response.Header.Get(headerName))
	}

	result, err := shortenRawBody(body, rules)
	if err != nil {
		return nil, err
	}

	return &ShortenedResponse{JSON: &result, StatusCode: response.StatusCode, Headers: resultHeader}, nil
}

func shortenRawBody(body []byte, rules map[string]string) (map[string]any, error) {
	parsedJson, err := oj.Parse(body)
	if err != nil {
		return map[string]any{}, ErrWhileShorteningServerResponse
	}

	result, err := shortenJSON(parsedJson, rules)
	if err != nil {
		return map[string]any{}, ErrWhileShorteningServerResponse
	}

	return result, nil
}

func shortenJSON(json any, rules map[string]string) (map[string]any, error) {
	result := make(map[string]any)
	for rule_k, rule_v := range rules {
		expr, err := jp.ParseString(rule_v)
		if err != nil {
			return map[string]any{}, ErrWhileShorteningServerResponse
		}
		parsed := expr.Get(json)
		result[rule_k] = parsed
	}
	return result, nil
}
