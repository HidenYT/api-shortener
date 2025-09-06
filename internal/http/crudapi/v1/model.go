package http

import (
	db_model "github.com/HidenYT/api-shortener/internal/storage/db-model/api"
)

type OutgoingRequestConfigRequest struct {
	Url            string `json:"url" validate:"required,http_url"`
	Method         string `json:"method" validate:"required"`
	Body           string `json:"body"`
	ShortenedAPIID uint   `json:"shortened_api_id" validate:"required"`
}

func outgoingRequestConfigRequestToDBModel(request *OutgoingRequestConfigRequest) *db_model.OutgoingRequestConfig {
	return &db_model.OutgoingRequestConfig{
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

func outgoingRequestHeaderRequestToDBModel(request *OutgoingRequestHeaderRequest) *db_model.OutgoingRequestHeader {
	return &db_model.OutgoingRequestHeader{
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

func outgoingRequestParamRequestToDBModel(request *OutgoingRequestParamRequest) *db_model.OutgoingRequestParam {
	return &db_model.OutgoingRequestParam{
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

func shorteningRuleRequestToDBModel(request *ShorteningRuleRequest) *db_model.ShorteningRule {
	return &db_model.ShorteningRule{
		FieldName:       request.FieldName,
		FieldValueQuery: request.FieldValueQuery,
		ShortenedAPIID:  request.ShortenedAPIID,
	}
}
