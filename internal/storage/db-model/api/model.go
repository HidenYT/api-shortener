package db_model

import (
	"gorm.io/gorm"
)

type ShortenedAPI struct {
	gorm.Model

	ShorteningRules []ShorteningRule `gorm:"constraint:OnDelete:CASCADE;"`
}

type OutgoingRequestConfig struct {
	gorm.Model

	Url    string `json:"url" validate:"required,http_url"`
	Method string `json:"method" validate:"required"`

	Headers []OutgoingRequestHeader `gorm:"constraint:OnDelete:CASCADE;"`
	Params  []OutgoingRequestParam  `gorm:"constraint:OnDelete:CASCADE;"`
	Body    string

	ShortenedAPIID uint          `validate:"required"`
	ShortenedAPI   *ShortenedAPI `gorm:"constraint:OnDelete:CASCADE;"`
}

type OutgoingRequestHeader struct {
	gorm.Model

	Name  string `validate:"required"`
	Value string `validate:"required"`

	OutgoingRequestConfigID uint `validate:"required"`
}

func (o *OutgoingRequestHeader) GetName() string {
	return o.Name
}

func (o *OutgoingRequestHeader) GetID() uint {
	return o.ID
}

func (o *OutgoingRequestHeader) SetID(id uint) {
	o.ID = id
}

type OutgoingRequestParam struct {
	gorm.Model

	Name  string `validate:"required"`
	Value string `validate:"required"`

	OutgoingRequestConfigID uint `validate:"required"`
}

func (o *OutgoingRequestParam) GetName() string {
	return o.Name
}

func (o *OutgoingRequestParam) GetID() uint {
	return o.ID
}

func (o *OutgoingRequestParam) SetID(id uint) {
	o.ID = id
}

type ShorteningRule struct {
	gorm.Model

	FieldName       string `validate:"required"`
	FieldValueQuery string `validate:"required,jsonpath-query"`

	ShortenedAPIID uint `validate:"required"`
}

func (o *ShorteningRule) GetName() string {
	return o.FieldName
}

func (o *ShorteningRule) GetID() uint {
	return o.ID
}

func (o *ShorteningRule) SetID(id uint) {
	o.ID = id
}
