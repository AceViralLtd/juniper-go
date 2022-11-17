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
		Expected map[string]string
		Data     t_Rules
	}{
		{"No Errors", map[string]string{}, t_Rules{"Frank", "frank@email.com", 25, 1, "FF"}},
		{"All Errors", map[string]string{
			"name":     "required",
			"email":    "required",
			"age":      "required",
			"Gender":   "required",
			"Initials": "required",
		}, t_Rules{}},
		{"One Error", map[string]string{"email": "required"}, t_Rules{"George", "", 20, 2, "GG"}},
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
