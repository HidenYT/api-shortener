package http

type ShortenedAPIDTO struct {
	ID                    uint                     `json:"id"`
	OutgoingRequestConfig OutgoingRequestConfigDTO `json:"outgoingRequestConfig"`
	ShorteningRules       []*ShorteningRuleDTO     `json:"shorteningRules"`
}

type OutgoingRequestConfigDTO struct {
	ID uint `json:"id"`

	Url    string `json:"url" validate:"required,http_url"`
	Method string `json:"method" validate:"required"`

	Headers []*OutgoingRequestHeaderDTO `json:"headers"`
	Params  []*OutgoingRequestParamDTO  `json:"params"`
	Body    string                      `json:"body"`
}

type OutgoingRequestHeaderDTO struct {
	ID uint `json:"id"`

	Name  string `json:"name" validate:"required"`
	Value string `json:"value" validate:"required"`
}

type OutgoingRequestParamDTO struct {
	ID uint `json:"id"`

	Name  string `json:"name" validate:"required"`
	Value string `json:"value" validate:"required"`
}

type ShorteningRuleDTO struct {
	ID uint `json:"id"`

	FieldName       string `json:"fieldName" validate:"required"`
	FieldValueQuery string `json:"fieldValueQuery" validate:"required,jsonpath-query"`
}
