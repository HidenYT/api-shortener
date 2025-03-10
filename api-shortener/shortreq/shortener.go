package shortreq

import (
	"fmt"

	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/oj"
)

type IResponseShortener interface {
	Shorten(responseBody []byte, rules map[string]string) (map[string]any, error)
}

type JSONResponseShortener struct{}

type ShorteningError struct {
	err error
}

func (e *ShorteningError) Error() string {
	return fmt.Sprintf("Error while shortening response: %s", e.err.Error())
}

func (shortener *JSONResponseShortener) Shorten(body []byte, rules map[string]string) (map[string]any, error) {
	parsedJson, err := oj.Parse(body)
	if err != nil {
		return map[string]any{}, &ShorteningError{err: err}
	}

	result, err := shortener.shortenWithRules(parsedJson, rules)
	if err != nil {
		return map[string]any{}, &ShorteningError{err: err}
	}

	return result, nil
}

func (shortener *JSONResponseShortener) shortenWithRules(json any, rules map[string]string) (map[string]any, error) {
	result := make(map[string]any)
	for rule_k, rule_v := range rules {
		expr, err := jp.ParseString(rule_v)
		if err != nil {
			return map[string]any{}, &ShorteningError{err: err}
		}
		parsed := expr.Get(json)
		result[rule_k] = parsed
	}
	return result, nil
}

func NewJsonResponseShortener() *JSONResponseShortener {
	return new(JSONResponseShortener)
}
