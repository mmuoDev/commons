package httputils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

//Get makes a HTTP GET call
func Get(URL string, w http.ResponseWriter) {
	client := http.Client{}
	request, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		ServeInternalError(err, w)
	}
	resp, err := client.Do(request)
	if err != nil {
		ServeInternalError(err, w)
	}
	defer resp.Body.Close()
	ServeJSON(resp, w)
}

//Post makes a HTTP post call
func Post(URL string, r interface{}, w http.ResponseWriter) {
	client := http.Client{}
	var body io.Reader
	bb, err := json.Marshal(r)
	if err != nil {
		ServeInternalError(err, w)
	}
	body = bytes.NewReader(bb)
	request, err := http.NewRequest(http.MethodPost, URL, body)
	if err != nil {
		ServeInternalError(err, w)
	}
	resp, err := client.Do(request)
	if err != nil {
		ServeInternalError(err, w)
	}
	defer resp.Body.Close()
	ServeJSON(resp, w)
}
