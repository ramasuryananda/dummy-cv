package app

import (
	httpwrapper "github.com/ramasuryananda/dummy-cv/internal/middleware/http-wrapper"
	"github.com/ramasuryananda/dummy-cv/internal/middleware/log_request"
	panichandler "github.com/ramasuryananda/dummy-cv/internal/middleware/panic_handler"
)

// Middleware types of middleware layer.
type Middleware struct {
	LogRequest   *log_request.Middleware
	PanicHandler *panichandler.Middleware
	HttpWrapper  *httpwrapper.Middleware
}

// NewMiddleware initializes middleware
func NewMiddleware() *Middleware {
	return &Middleware{
		LogRequest:   log_request.New(),
		PanicHandler: panichandler.New(),
		HttpWrapper:  httpwrapper.New(),
	}
}
