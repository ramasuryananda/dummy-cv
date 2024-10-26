package constant

import (
	"errors"
	"net/http"

	"github.com/ramasuryananda/dummy-cv/internal/dto/general"
)

var (
	ErrorDatabaseNotFound = errors.New("data not found in database")

	ResponseSuccess             = general.ResponseCode{Code: "00", Description: "Success", Status: http.StatusOK}
	ResponseValidationError     = general.ResponseCode{Code: "VE", Description: "Validation Error", Status: http.StatusBadRequest}
	ResponseErrorNotFound       = general.ResponseCode{Code: "NF", Description: "Data Not Found", Status: http.StatusNotFound}
	ResponseInternalServerError = general.ResponseCode{Code: "IS", Description: "Something went wrong, try again later", Status: http.StatusInternalServerError}
	ResponseBadRequest          = general.ResponseCode{Code: "BR", Description: "Bad Request", Status: http.StatusBadRequest}
)
