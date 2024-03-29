package httputils

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/mmuoDev/commons/time"
	"github.com/mmuoDev/commons/errors"
	pkgErr "github.com/pkg/errors"
	"github.com/gorilla/schema"
)

//ServeJSON returns a JSON response for a http request
func ServeJSON(res interface{}, w http.ResponseWriter) {

	bb, err := json.Marshal(res)
	if err != nil {
		ServeInternalError(err, w)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write(bb)
}

//ServeInternalError returns a 500 response for a http request
func ServeInternalError(err error, w http.ResponseWriter) {
	var errDTO errors.Error
	errDTO, ok := err.(errors.Error)

	if !ok {
		errDTO = errors.Error{
			Message: errors.ErrorMessage(err.Error()),
		}
	}

	bb, err := json.Marshal(errDTO)
	if err != nil {
		ServeInternalError(err, w)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(bb)
}

// ServeError is a generic error function that serves a custom error depending on params
func ServeError(err error, w http.ResponseWriter) {
	if errors.IsNotFoundError(err) {
		ServeNotFoundError(err, w)
		return
	}

	if errors.IsBadRequestError(err) {
		ServeBadRequestError(err, w)
		return
	}

	if errors.IsConflictError(err) {
		ServeConflictError(err, w)
		return
	}

	ServeInternalError(err, w)
}

// ServeNotFoundError returns a 404 response for an http request
func ServeNotFoundError(err error, w http.ResponseWriter) {
	var errDTO errors.Error
	errDTO, ok := err.(errors.Error)

	if !ok {
		errDTO = errors.Error{
			Message: errors.ErrorMessage(err.Error()),
		}
	}

	bb, err := json.Marshal(errDTO)
	if err != nil {
		ServeNotFoundError(err, w)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusNotFound)
	w.Write(bb)
}

// ServeConflictError returns a 409 response for an http request
func ServeConflictError(err error, w http.ResponseWriter) {
	var errDTO errors.Error
	errDTO, ok := err.(errors.Error)

	if !ok {
		errDTO = errors.Error{
			Message: errors.ErrorMessage(err.Error()),
		}
	}

	bb, err := json.Marshal(errDTO)
	if err != nil {
		ServeInternalError(err, w)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusConflict)
	w.Write(bb)
}

// ServeBadRequestError returns a 400 response for an http request
func ServeBadRequestError(err error, w http.ResponseWriter) {

	var errDTO errors.Error
	errDTO, ok := err.(errors.Error)

	if !ok {
		errDTO = errors.Error{
			Message: errors.ErrorMessage(err.Error()),
		}
	}

	bb, err := json.Marshal(errDTO)
	if err != nil {
		ServeInternalError(err, w)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(bb)
}

// JSONToDTO decodes an http request JSON body to a data transfer object
func JSONToDTO(DTO interface{}, w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&DTO); err != nil {
		ServeInternalError(err, w)
	}
}

// GetQueryParam retreives a query param from an incoming http request if no param exists returns empty string
func GetQueryParam(key string, r *http.Request) string {
	param, ok := r.URL.Query()[key]
	if !ok {
		return ""
	}
	return param[0]
}

// FileToStruct reads a json file to a struct.. a reader for the file bytes is also returned
func FileToStruct(filepath string, s interface{}) io.Reader {
	bb, _ := ioutil.ReadFile(filepath)
	json.Unmarshal(bb, s)
	return bytes.NewReader(bb)
}

var decoder = schema.NewDecoder()

// GetQueryParams maps the query params from an http request into an interface
func GetQueryParams(value interface{}, r *http.Request) error {
	// decoder lookup for values on the json tag, instead of the default schema tag
	decoder.SetAliasTag("json")

	var globalErr error

	// Decoder Register for custom type ISO8601
	decoder.RegisterConverter(time.ISO8601{}, func(input string) reflect.Value {
		ISOTime, errISO := time.NewISO8601(input)

		if errISO != nil {
			globalErr = pkgErr.Wrapf(errISO, "handler - invalid iso time provided")
			return reflect.ValueOf(time.ISO8601{})
		}

		return reflect.ValueOf(ISOTime)
	})

	// Decoder Register for custom type Epoch
	decoder.RegisterConverter(time.Epoch(0), func(input string) reflect.Value {
		ISOTime, errISO := time.NewISO8601(input)

		if errISO != nil {
			globalErr = pkgErr.Wrapf(errISO, "handler - invalid iso time provided")
			return reflect.ValueOf(time.ISO8601{}.ToEpoch())
		}

		return reflect.ValueOf(ISOTime.ToEpoch())
	})

	if err := decoder.Decode(value, r.URL.Query()); err != nil {
		return pkgErr.Wrapf(err, "handler - failed to decode query params")
	}

	if globalErr != nil {
		return globalErr
	}

	return nil
}
