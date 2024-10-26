package constant

import (
	"errors"
	"net/http"

	"github.com/ramasuryananda/dummy-cv/internal/dto/general"
)

var (
	ErrorDatabaseNotFound = errors.New("data not found in database")

	ResponseSuccess       = general.ResponseCode{Code: "00", Description: "Success", Status: http.StatusOK}
	ResponseErrorNotFound = general.ResponseCode{Code: "NF", Description: "Data Not Found", Status: http.StatusNotFound}
)
