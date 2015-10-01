package handler

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/labstack/echo"
	"github.com/shwoodard/jsonapi"
)

// decode is a helper function for decoding json bodies into objects. decode
// also closes the request body.
func decode(r *http.Request, v interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}

// encodeJSON is a helper function for returning objects in a JSON API spec format.
func encodeJSON(c *echo.Context, code int, v interface{}) error {
	c.Response().Header().Set("Content-Type", "application/vnd.api+json")
	c.Response().WriteHeader(code)

	s := reflect.ValueOf(v)
	if s.Kind() != reflect.Slice {
		return jsonapi.MarshalOnePayload(c.Response(), v)
	}

	ret := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}
	return jsonapi.MarshalManyPayload(c.Response(), ret)
}

// newError is a helper function that will convert the erorr in a JSON response.
func newError(code int, desc string) HTTPError {
	return HTTPError{code, desc}
}

type HTTPError struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

// Error implements the error interface.
func (err HTTPError) Error() string {
	return err.Description
}
