package juniper

import (
	"reflect"
	"testing"

	"github.com/go-playground/validator/v10"
)

type t_Rules struct {
	Name     string `validate:"required" form:"name"`
	Email    string `validate:"required" json:"email"`
	Age      int    `validate:"required" from:"-" json:"age"`
	Gender   int    `validate:"required" form:"-" json:"-"`
	Initials string `validate:"required"`
}

func TestErrorExtraction(t *testing.T) {
	testCases := []struct {
		TestName string
		Expected []string
		Data     t_Rules
	}{
		{"No Errors", nil, t_Rules{"Frank", "frank@email.com", 25, 1, "FF"}},
		{"All Errors", []string{"name", "email", "age", "Gender", "Initials"}, t_Rules{}},
		{"One Error", []string{"email"}, t_Rules{"George", "", 20, 2, "GG"}},
	}

	validate := validator.New()
	validate.RegisterTagNameFunc(ValidationFieldExtractor)

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			err := validate.Struct(testCase.Data)
			errorFields := errorFIelds(err)

			if !reflect.DeepEqual(testCase.Expected, errorFields) {
				t.Fatalf("Expected: %#v, Found: %#v", testCase.Expected, errorFields)
			}
		})
	}
}
