package binding

import (
	"go.gitbaidu.com/gopkg/utility/check"
	"gopkg.in/go-playground/validator.v9"
)

func IsNoHTML(fl validator.FieldLevel) bool {
	return !check.IsHTML(fl.Field().String())
}

func IsColumn(fl validator.FieldLevel) bool {
	return check.IsColumn(fl.Field().String())
}
