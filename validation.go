package juniper

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type inputValidator struct {
	validator *validator.Validate
}

func (val *inputValidator) Validate(i interface{}) error {
	return val.validator.Struct(i)
}

// InitializeValidator with the field extractor procedure
func InitializeValidator(router *echo.Echo) {
	v := &inputValidator{validator.New()}
	v.validator.RegisterTagNameFunc(ValidationFieldExtractor)

	router.Validator = v
}

// ValidationFieldExtractor will extract the most appripriate field name it can
func ValidationFieldExtractor(field reflect.StructField) string {
	formName := strings.SplitN(field.Tag.Get("form"), ",", 2)[0]
	if formName != "-" && formName != "" {
		return formName
	}

	jsonName := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
	if jsonName != "-" && jsonName != "" {
		return jsonName
	}

	return field.Name
}

// errorFIelds wil extract the field names from validation errors
//
// this will be the JSON field names, not the go struct field names
func errorFIelds(errors interface{}) []string {
	var fields []string

	switch val := errors.(type) {
	case validator.ValidationErrors:
		for _, err := range val {
			fields = append(fields, err.Field())
		}
	}

	return fields
}
