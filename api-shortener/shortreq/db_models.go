package shortreq

import (
	"time"

	"gorm.io/gorm"
)

type ShortenedAPI struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	ConfigID uint                   `json:"-"`
	Config   *OutgoingRequestConfig `json:"config" validate:"required"`

	ShorteningRules []ShorteningRule `json:"rules" validate:"required,dive"`
}

type OutgoingRequestConfig struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	Url    string `json:"url" validate:"required,http_url"`
	Method string `json:"method" validate:"required"`

	Headers []*OutgoingRequestHeader `json:"headers" validate:"dive"`
	Params  []*OutgoingRequestParam  `json:"params" validate:"dive"`
	Body    string                   `json:"body"`
}

type OutgoingRequestHeader struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	Name  string `json:"name" validate:"required"`
	Value string `json:"value" validate:"required"`

	OutgoingRequestConfigID uint                   `json:"-"`
	OutgoingRequestConfig   *OutgoingRequestConfig `json:"-"`
}

type OutgoingRequestParam struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	Name  string `json:"name" validate:"required"`
	Value string `json:"value" validate:"required"`

	OutgoingRequestConfigID uint                   `json:"-"`
	OutgoingRequestConfig   *OutgoingRequestConfig `json:"-"`
}

type ShorteningRule struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	FieldName       string `json:"field_name" validate:"required"`
	FieldValueQuery string `json:"field_value_query" validate:"required,jsonpath-query"`

	ShortenedAPIID uint          `json:"-"`
	ShortenedAPI   *ShortenedAPI `json:"-"`
}
