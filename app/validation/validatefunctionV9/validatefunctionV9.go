package validatefunctionV9

import "gopkg.in/go-playground/validator.v9"

func AwsomeValidate(fl validator.FieldLevel) bool {
	if fl.Field().String() != "awsome" {
		return false
	}
	return true
}
