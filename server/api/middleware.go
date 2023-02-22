package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"renthome/boiler"

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
	Code    string `json:"code"`
	Message string `json:"message"`
}

// WithError middleware handles the HTTP error response
func WithError(next func(w http.ResponseWriter, r *http.Request) (int, error)) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		contents, err := io.ReadAll(r.Body)
		if err != nil {
			terror.Echo(err)
		}
		r.Body = io.NopCloser(bytes.NewReader(contents))
		defer r.Body.Close()

		code, err := next(w, r)
		if err != nil {
			terror.Echo(err)
			errObj := &HTTPErrorResponse{
				Code:    fmt.Sprintf("%d", code),
				Message: err.Error(),
			}

			var tErr *terror.TError
			if errors.As(err, &tErr) {
				errObj.Message = tErr.Message

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
			}

			jsonErr, err := json.Marshal(errObj)
			if err != nil {
				terror.Echo(err)
				http.Error(w, `{"code":"00001","message":"JSON failed, please contact IT."}`, code)
				return
			}

			http.Error(w, string(jsonErr), code)
			return
		}
	}

	return fn
}

// WithUser checks for authenticated user
func WithUser(api *APIController, next func(w http.ResponseWriter, r *http.Request, user *boiler.User) (int, error)) func(w http.ResponseWriter, r *http.Request) (int, error) {
	fn := func(w http.ResponseWriter, r *http.Request) (int, error) {
		user, err := GetUserFromToken(api, r)
		if err != nil {
			return http.StatusUnauthorized, terror.Error(err, ErrUnauthorised)
		}

		if user != nil {
			return next(w, r, user)
		}

		return http.StatusUnauthorized, terror.Error(fmt.Errorf("unathorised"), ErrUnauthorised)

	}
	return fn
}
