package errors

import (
	"fmt"
)

// FieldName is the name of the field that failed validation
type FieldName string

// ErrMessage is the message describing the field validation failure
type ErrMessage interface{}

// FieldErrors is a key value collection of the names of all failed fields
// and their corresponding failure messages
type FieldErrors map[FieldName]ErrMessage

// Error adds the error behavior to ValidationError
func (v FieldErrors) Error() string {
	return fmt.Sprintln("Field errors exsits")

}

// IsFieldValidationError returns trye if the error is of type fieldValidationFailure
func IsFieldValidationError(err error) bool {
	_, ok := err.(FieldErrors)
	return ok
}

// RetrieveFieldName returns the underlying field name for f
func RetrieveFieldName(f FieldName) string {
	return string(f)
}

// RetrieveMessage returns the underlying message for m
func RetrieveMessage(m ErrMessage) string {
	return m.(string)
}
