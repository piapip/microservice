package data

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator"
)

// ValidationError wraps the validators FieldError so we do not expose this to out code
type ValidationError struct {
	validator.FieldError
}

func (v ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error: Field validation for '%s' failed on the '%s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

// ValidationErrors is a collection of ValidationError
type ValidationErrors []ValidationError

// Errors converts the slice into a string slice
func (v ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

// Validation contains the validator settings and cache
type Validation struct {
	validate *validator.Validate
}

// NewValidation creates new validation type
func NewValidation() *Validation {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)

	return &Validation{validate}
}

// Validate the item
// for more detail the returned error can be cast into a
// validator.ValidationErrors collection
//
// if ve, ok := err.(validator.ValidationErrors); ok {
// 	fmt.Println(ve.Namespace())
// 	fmt.Println(ve.Field())
// 	fmt.Println(ve.StructNamespace())
// 	fmt.Println(ve.StructField())
// 	fmt.Println(ve.Tag())
// 	fmt.Println(ve.ActualTag())
// 	fmt.Println(ve.Kind())
// 	fmt.Println(ve.Type())
// 	fmt.Println(ve.Value())
// 	fmt.Println(ve.Param())
// 	fmt.Println()
// }
func (v *Validation) Validate(i interface{}) ValidationErrors {
	// .Struct(i) validates a structs exposed fields
	// .(validator.ValidationErrors) will type cast the error returned from .Struct(i) to type ValidationErrors

	// WARNING!!! YOU CAN'T CAST .(validator.ValidationErrors) this right away.
	// This error is seen quite often in nodejs, when you can't cast struct to a nil.
	// You'll have to check for nil first, otherwise it will shoot error. Cost me an hour.
	// errs := v.validate.Struct(i).(validator.ValidationErrors)
	errs := v.validate.Struct(i)
	if errs != nil {
		var returnErrs ValidationErrors
		for _, err := range errs.(validator.ValidationErrors) {
			// cast the FieldError into our ValidationError and append to the slice
			validationError := ValidationError{err.(validator.FieldError)}
			returnErrs = append(returnErrs, validationError)
		}

		return returnErrs
	}

	return nil
}

// validateSKU
func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}

	return true
}
