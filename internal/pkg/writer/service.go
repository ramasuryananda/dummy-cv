package writer

import (
	"errors"
	"net/http"

	"github.com/ramasuryananda/dummy-cv/internal/pkg/exception"
)

func APIResponse(message string, status bool, data interface{}) Response {
	jsonResponse := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}

	return jsonResponse
}

func APIErrorResponse(message string, err error) (response Response, code int) {

	code = http.StatusInternalServerError
	var exception *exception.Exception
	if errors.As(err, &exception) {
		message = err.Error()
		code = http.StatusBadRequest
	}

	jsonResponse := Response{
		Status:  false,
		Message: message,
	}

	return jsonResponse, code
}

func APIValidationResponse(message string, status bool, data interface{}, errors interface{}) ValidationResponse {
	jsonResponse := ValidationResponse{
		Status:  status,
		Message: message,
		Data:    data,
		Errors:  errors,
	}

	return jsonResponse
}
