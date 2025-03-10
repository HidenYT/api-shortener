package restapi

type OutgoingRequestConfigRequest struct {
	Url            string `json:"url" validate:"required,http_url"`
	Method         string `json:"method" validate:"required"`
	Body           string `json:"body"`
	ShortenedAPIID uint   `json:"shortened_api_id" validate:"required"`
}

func outgoingRequestConfigRequestToDBModel(request *OutgoingRequestConfigRequest) *OutgoingRequestConfig {
	return &OutgoingRequestConfig{
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

func outgoingRequestHeaderRequestToDBModel(request *OutgoingRequestHeaderRequest) *OutgoingRequestHeader {
	return &OutgoingRequestHeader{
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

func outgoingRequestParamRequestToDBModel(request *OutgoingRequestParamRequest) *OutgoingRequestParam {
	return &OutgoingRequestParam{
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

func shorteningRuleRequestToDBModel(request *ShorteningRuleRequest) *ShorteningRule {
	return &ShorteningRule{
		FieldName:       request.FieldName,
		FieldValueQuery: request.FieldValueQuery,
		ShortenedAPIID:  request.ShortenedAPIID,
	}
}
