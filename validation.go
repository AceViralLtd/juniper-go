package juniper

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// InitializeValidator with the field extractor procedure
func InitializeValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(ValidationFieldExtractor)
	}
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
func errorFIelds(errors interface{}) map[string]string {
	fields := make(map[string]string)

	switch val := errors.(type) {
	case validator.ValidationErrors:
		for _, err := range val {
			fields[err.Field()] = err.Tag()
		}
	}

	return fields
}
