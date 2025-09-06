package http

import api_dao "github.com/HidenYT/api-shortener/internal/storage/dao"

type OutgoingRequestConfigRequest struct {
	Url            string `json:"url" validate:"required,http_url"`
	Method         string `json:"method" validate:"required"`
	Body           string `json:"body"`
	ShortenedAPIID uint   `json:"shortened_api_id" validate:"required"`
}

func outgoingRequestConfigRequestToDBModel(request *OutgoingRequestConfigRequest) *api_dao.OutgoingRequestConfig {
	return &api_dao.OutgoingRequestConfig{
		Url:            request.Url,
		Method:         request.Method,
		Body:           request.Body,
		ShortenedAPIID: request.ShortenedAPIID,
	}
}

type OutgoingRequestHeaderRequest struct {
	Name                    string `json:"name" validate:"required"`
	Value                   string `json:"value" validate:"required"`
	OutgoingRequestConfigID uint   `json:"outgoing_request_config_id" validate:"required"`
}

func outgoingRequestHeaderRequestToDBModel(request *OutgoingRequestHeaderRequest) *api_dao.OutgoingRequestHeader {
	return &api_dao.OutgoingRequestHeader{
		Name:                    request.Name,
		Value:                   request.Value,
		OutgoingRequestConfigID: request.OutgoingRequestConfigID,
	}
}

type OutgoingRequestParamRequest struct {
	Name                    string `json:"name" validate:"required"`
	Value                   string `json:"value" validate:"required"`
	OutgoingRequestConfigID uint   `json:"outgoing_request_config_id" validate:"required"`
}

func outgoingRequestParamRequestToDBModel(request *OutgoingRequestParamRequest) *api_dao.OutgoingRequestParam {
	return &api_dao.OutgoingRequestParam{
		Name:                    request.Name,
		Value:                   request.Value,
		OutgoingRequestConfigID: request.OutgoingRequestConfigID,
	}
}

type ShorteningRuleRequest struct {
	FieldName       string `json:"field_name" validate:"required"`
	FieldValueQuery string `json:"field_value_query" validate:"required,jsonpath-query"`
	ShortenedAPIID  uint   `json:"shortened_api_id" validate:"required"`
}

func shorteningRuleRequestToDBModel(request *ShorteningRuleRequest) *api_dao.ShorteningRule {
	return &api_dao.ShorteningRule{
		FieldName:       request.FieldName,
		FieldValueQuery: request.FieldValueQuery,
		ShortenedAPIID:  request.ShortenedAPIID,
	}
}
