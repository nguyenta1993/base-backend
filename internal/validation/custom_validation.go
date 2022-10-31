package validation

import (
	"github.com/go-playground/validator/v10"
)

const Tag = "age_not_negative"

var AgeNotNegative validator.Func = func(fl validator.FieldLevel) bool {
	return fl.Field().Int() > 0
}
