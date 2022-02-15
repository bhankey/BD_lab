package clienterror

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AppError struct {
	Err              error  `json:"-"`
	Message          string `json:"message,omitempty"`
	DeveloperMessage string `json:"developer_message,omitempty"`
	Code             string `json:"code,omitempty"`
	HTTPCode         int    `json:"-"`
}

var ErrStatusCode = fmt.Errorf("wrong http status code")

func NewAppError(message, code, developerMessage string) *AppError {
	return &AppError{
		Err:              fmt.Errorf(message),
		Code:             code,
		Message:          message,
		DeveloperMessage: developerMessage,
	}
}

func (ae *AppError) Error() string {
	return ae.Err.Error()
}

func (ae *AppError) Unwrap() error {
	return ae.Err
}

func (ae *AppError) Marshal() ([]byte, error) {
	return json.Marshal(ae)
}

func (ae *AppError) SetHTTPCode(code int) error {
	if code < 100 || code > 600 {
		return ErrStatusCode
	}
	ae.HTTPCode = code

	return nil
}

func (ae *AppError) GetHTTPCode() int {
	if ae == nil || ae.HTTPCode == 0 {
		return http.StatusInternalServerError
	}

	return ae.HTTPCode
}
