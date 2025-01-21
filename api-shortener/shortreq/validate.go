package shortreq

import (
	"github.com/go-playground/validator/v10"
	"github.com/ohler55/ojg/jp"
)

func NewValidate() *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation("jsonpath-query", ValidateJSONPathQuery)
	return validate
}

func ValidateJSONPathQuery(fl validator.FieldLevel) bool {
	_, err := jp.ParseString(fl.Field().String())
	return err == nil
}
