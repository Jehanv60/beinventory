package util

import (
	"fmt"
	"regexp"

	"github.com/Jehanv60/exception"
	"github.com/Jehanv60/helper"
	"github.com/go-playground/validator/v10"
)

func ValidateAlphanumdash(fl validator.FieldLevel) bool {
	validation := fl.Field().String()
	simbol, err := regexp.Compile("^[-a-zA-Z0-9 _]*$")
	helper.PanicError(err)
	result := simbol.MatchString(validation)
	return result
}

func ErrValidateSelf(err error) {
	var errValTag []error
	if err != nil {
		for _, errVal := range err.(validator.ValidationErrors) {
			var errCatch error
			switch errVal.Tag() {
			case "alphanumdash":
				errCatch = fmt.Errorf("%s:Format Tidak Boleh Pakai Simbol", errVal.Field())
			case "email":
				errCatch = fmt.Errorf("%s:Format Harus Email", errVal.Field())
			case "required":
				errCatch = fmt.Errorf("%s:Tidak Boleh Kosong", errVal.Field())
			case "alphanum":
				errCatch = fmt.Errorf("%s:Tidak Boleh Spasi dan Simbol", errVal.Field())
			case "gte":
				errCatch = fmt.Errorf("%s:Angka Tidak Boleh Mines", errVal.Field())
			case "lte":
				errCatch = fmt.Errorf("%s:Angka Tidak Boleh Melebihi %s", errVal.Field(), errVal.Param())
			default:
				errCatch = fmt.Errorf("error Pada Field: %s Dan Err :%s", errVal.Field(), errVal)
			}
			errValTag = append(errValTag, errCatch)
		}
		panic(exception.NewValidateFound(errValTag))
	}
}
