package restapi

import (
	"time"

	"gorm.io/gorm"
)

type ShortenedAPI struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	ShorteningRules []ShorteningRule `json:"-"`
}

type OutgoingRequestConfig struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	Url    string `json:"url" validate:"required,http_url"`
	Method string `json:"method" validate:"required"`

	Headers []*OutgoingRequestHeader `json:"-"`
	Params  []*OutgoingRequestParam  `json:"-"`
	Body    string                   `json:"body"`

	ShortenedAPIID uint          `json:"shortened_api_id" validate:"required"`
	ShortenedAPI   *ShortenedAPI `json:"-" gorm:"constraint:OnDelete:CASCADE"`
}

type OutgoingRequestHeader struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	Name  string `json:"name" validate:"required"`
	Value string `json:"value" validate:"required"`

	OutgoingRequestConfigID uint                   `json:"outgoing_request_config_id" validate:"required"`
	OutgoingRequestConfig   *OutgoingRequestConfig `json:"-" gorm:"constraint:OnDelete:CASCADE"`
}

type OutgoingRequestParam struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	Name  string `json:"name" validate:"required"`
	Value string `json:"value" validate:"required"`

	OutgoingRequestConfigID uint                   `json:"outgoing_request_config_id" validate:"required"`
	OutgoingRequestConfig   *OutgoingRequestConfig `json:"-" gorm:"constraint:OnDelete:CASCADE"`
}

type ShorteningRule struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	FieldName       string `json:"field_name" validate:"required"`
	FieldValueQuery string `json:"field_value_query" validate:"required,jsonpath-query"`

	ShortenedAPIID uint          `json:"shortened_api_id" validate:"required"`
	ShortenedAPI   *ShortenedAPI `json:"-" gorm:"constraint:OnDelete:CASCADE"`
}
