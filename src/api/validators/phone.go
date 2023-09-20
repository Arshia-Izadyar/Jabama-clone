package validators

import (
	"github.com/Arshia-Izadyar/Jabama-clone/src/common"
	"github.com/go-playground/validator/v10"
)

func IranPhoneNumberValidator(fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(string)
	if !ok {
		return false
	}
	return common.ValidateNumber(value)
}
