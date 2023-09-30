package validators

import (
	"github.com/go-playground/validator/v10"
)

func RateValidator(fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(int)
	if !ok {
		return false
	}
	return value <= 10
}
