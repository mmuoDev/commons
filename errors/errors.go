package errors

import (
	"encoding/json"
	"fmt"

	pkgerrs "github.com/pkg/errors"
)

// ErrorCode is a machine friendly description of the error
type ErrorCode string

const (
	invalidReqBodyErrCode     = ErrorCode("invalid_request_body")
	invalidReqErrMessage      = ErrorMessage("One or more input values are invalid. Please re-renter valid values")
	repeatedDocumentErrCode   = ErrorCode("document_already_exists")
	invalidDocumentErrMessage = ErrorMessage("A document with the provided unique identifier already exists")
)

// ErrorMessage is a human friendly description of the error
type ErrorMessage string

// ErrorParams captures an additional key value pairs of information about an error
type ErrorParams map[string]interface{}

// NewErrorParams creates a new error params with a first entry
func NewErrorParams(k, v string) ErrorParams {
	pp := map[string]interface{}{}
	pp[k] = v
	return pp
}

// ToBadRequest converts error params to a Bad Request Error
func (p ErrorParams) ToBadRequest() Error {
	return BadRequestError(p)
}

// ToConflict converts error params to a Conflict Error
func (p ErrorParams) ToConflict() Error {
	return ConflictError(p)
}

// ConflictError creates a new conflict request error
func ConflictError(p ErrorParams) Error {
	return Error{
		Code:    repeatedDocumentErrCode,
		Message: invalidDocumentErrMessage,
		Params:  p,
	}
}

// IsConflictError checks if an error is a not conflict error
func IsConflictError(err error) bool {
	cause := pkgerrs.Cause(err)
	esusuErr, ok := cause.(Error)
	return ok && esusuErr.Code == repeatedDocumentErrCode
}

// An Error is a representation of an api error
// Errors include but not limited to domain errors, validation, authentication error etc
type Error struct {
	Code    ErrorCode    `json:"code,omitempty"`
	Message ErrorMessage `json:"message"`
	Params  ErrorParams  `json:"params,omitempty"`
}

func (e Error) Error() string {
	msg, _ := json.MarshalIndent(&e, "", "\t")
	return string(msg)
}

// BadRequestError creates a new bad request error
func BadRequestError(p ErrorParams) Error {
	return Error{
		Code:    invalidReqBodyErrCode,
		Message: invalidReqErrMessage,
		Params:  p,
	}
}

//FromFieldError creates an Error from field errors
func FromFieldError(fieldErr FieldErrors) Error {
	params := map[string]interface{}{}
	for k, v := range fieldErr {
		params[RetrieveFieldName(k)] = v
	}
	return BadRequestError(params)
}

// IsBadRequestError checks if an error is a not found error
func IsBadRequestError(err error) bool {
	cause := pkgerrs.Cause(err)
	esusuErr, ok := cause.(Error)
	return ok && esusuErr.Code == invalidReqBodyErrCode
}

type entityNotFoundError struct {
	message    string
	wrappedErr error
}

func (e entityNotFoundError) Error() string {
	return fmt.Sprintf("%s err=%s", e.message, e.wrappedErr.Error())
}

// NotFound creates a new NotFoundError
func NotFound(err error, msg string) error {
	return entityNotFoundError{
		message:    msg,
		wrappedErr: err,
	}
}

// IsNotFoundError checks if an error is a not found error
func IsNotFoundError(err error) bool {
	cause := pkgerrs.Cause(err)
	_, ok := cause.(entityNotFoundError)
	return ok
}
