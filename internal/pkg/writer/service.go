package writer

import "github.com/ramasuryananda/dummy-cv/internal/constant"

func APIResponse(code string, message string, data interface{}) Response {
	jsonResponse := Response{
		Code:    code,
		Message: message,
		Data:    data,
	}

	return jsonResponse
}

func APIErrorResponse(responseCode string, message string, err error) (response Response) {
	jsonResponse := Response{
		Code:    responseCode,
		Message: message,
	}

	return jsonResponse
}

func APIValidationResponse(data interface{}, errors interface{}) ValidationResponse {
	jsonResponse := ValidationResponse{
		Code:    constant.ResponseValidationError.Code,
		Message: constant.ResponseErrorNotFound.Description,
		Data:    data,
		Errors:  errors,
	}

	return jsonResponse
}
