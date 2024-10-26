package translator

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator"
)

func TranslateError(err error, s interface{}) map[string]string {
	apiErrors := make(map[string]string)

	// Mapping translator errors
	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				apiErrors[err.Field()] = fmt.Sprintf("%s is required", getTagName(reflect.TypeOf(s), err.Field()))
			case "date":
				apiErrors[err.Field()] = fmt.Sprintf("%s invalid date format, the date format should be DD/MM/YYYY", getTagName(reflect.TypeOf(s), err.Field()))
			case "oneof":
				apiErrors[err.Field()] = fmt.Sprintf("%s is unknown", getTagName(reflect.TypeOf(s), err.Field()))
			case "numeric":
				apiErrors[err.Field()] = fmt.Sprintf("%s must be numeric", getTagName(reflect.TypeOf(s), err.Field()))
			case "max":
				apiErrors[err.Field()] = fmt.Sprintf("%s can't be greater than %s", getTagName(reflect.TypeOf(s), err.Field()), err.Param())
			case "unique":
				apiErrors[err.Field()] = fmt.Sprintf("%s must be unique", getTagName(reflect.TypeOf(s), err.Field()))
			case "min":
				apiErrors[err.Field()] = fmt.Sprintf("%s minimum of %s", getTagName(reflect.TypeOf(s), err.Field()), err.Param())

			}
		}
	}

	return apiErrors
}

func getTagName(t reflect.Type, fieldName string) string {
	field, _ := t.FieldByName(fieldName)
	name := field.Tag.Get("name")

	if name == "" {
		name = fieldName
	}

	return name
}
