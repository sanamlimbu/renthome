package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/ninja-software/terror/v2"
)

type ErrorMessage string

const (
	Unauthorized        ErrorMessage = "Unauthorized - Please log in and try again."
	Forbidden           ErrorMessage = "Forbidden - You do not have permission to perform this action."
	InternalServerError ErrorMessage = "Internal Server Error - Please try again in a few minutes."
	BadRequest          ErrorMessage = "Bad Request - Please try again with correct inputs."
)

func (errMsg ErrorMessage) String() string {
	return string(errMsg)
}

type HTTPErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// WithError middleware handles the HTTP error response
func WithError(next func(w http.ResponseWriter, r *http.Request) (int, error)) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		contents, err := ioutil.ReadAll(r.Body)
		r.Body = ioutil.NopCloser(bytes.NewReader(contents))
		defer r.Body.Close()

		code, err := next(w, r)
		if err != nil {
			terror.Echo(err)
			errObj := &HTTPErrorResponse{
				Code:    code,
				Message: err.Error(),
			}

			var tErr *terror.TError
			if errors.As(err, tErr) {
				errObj.Message = tErr.Message
			}

			// set generic message if friendly message not set
			if tErr.Error() == tErr.Message {
				if code == 500 {
					errObj.Message = InternalServerError.String()
				}

				if code == 403 {
					errObj.Message = Forbidden.String()
				}

				if code == 401 {
					errObj.Message = Unauthorized.String()
				}

				if code == 400 {
					errObj.Message = BadRequest.String()
				}
			}

			jsonErr, err := json.Marshal(errObj)
			if err != nil {
				terror.Echo(err)
				http.Error(w, `{"code": "00001",:message:"JSON failed, please contact IT."}`, code)
				return
			}

			http.Error(w, string(jsonErr), code)
			return
		}
	}

	return fn
}
