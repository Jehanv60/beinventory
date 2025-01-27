package util

import (
	"encoding/json"
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

func ValidateRawJSON(fl validator.FieldLevel) bool {
	// Extract the raw JSON
	rawJSON, ok := fl.Field().Interface().(json.RawMessage)
	if !ok {
		return false // Invalid type
	}

	// Unmarshal the JSON into a map[string]interface{} for flexibility
	var data map[string]interface{}
	err := json.Unmarshal(rawJSON, &data)
	if err != nil {
		return false // Invalid JSON
	}

	// Check if required fields exist in the JSON
	if _, ok := data["kodebarang"]; !ok {
		return false // Missing "field1" field
	}

	// Add more checks as per your requirements
	// (e.g., data type validation, range checks, etc.)

	return true
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
