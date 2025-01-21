package shortreq

import (
	"net/http"
)

type OutgoingRequestConfigModel struct {
	Url     string `validate:"required,http_url"`
	Method  string
	Headers http.Header
	Params  map[string]string
	Body    string
}

type IOutgoingRequestConfigResolver interface {
	GetRequestConfigModel(api *ShortenedAPI) *OutgoingRequestConfigModel
}

type OutgoingRequestResolver struct {
	dao IShortenedAPIDAO
}

func (resolver *OutgoingRequestResolver) GetRequestConfigModel(api *ShortenedAPI) *OutgoingRequestConfigModel {
	return &OutgoingRequestConfigModel{
		Url:     api.Config.Url,
		Method:  api.Config.Method,
		Headers: resolver.getHeadersFromDBHeaders(api.Config.Headers),
		Params:  resolver.getParamsFromDBParams(api.Config.Params),
		Body:    api.Config.Body,
	}
}

func (*OutgoingRequestResolver) getHeadersFromDBHeaders(dbHeaders []*OutgoingRequestHeader) http.Header {
	httpHeader := make(http.Header)
	for _, header := range dbHeaders {
		httpHeader.Add(header.Name, header.Value)
	}
	return httpHeader
}

func (*OutgoingRequestResolver) getParamsFromDBParams(dbParams []*OutgoingRequestParam) map[string]string {
	httpParams := make(map[string]string)
	for _, param := range dbParams {
		httpParams[param.Name] = param.Value
	}
	return httpParams
}

func NewOutgoingRequestResolver(apiDAO IShortenedAPIDAO) IOutgoingRequestConfigResolver {
	return &OutgoingRequestResolver{dao: apiDAO}
}
