package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	app "github.com/ramasuryananda/dummy-cv/internal/app"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/config"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/customvalidator"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	appLogger "github.com/ramasuryananda/dummy-cv/internal/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	timeoutServer = 60
	port          = 8080
)

type Server struct {
	handler    *app.Handlers
	http       *http.Server
	middleware *app.Middleware
}

func NewHTTP(ctx context.Context) *Server {

	appLogger.Init(config.Get("APP_NAME"))

	db, err := mysqlConnect()
	if err != nil {
		panic(err)
	}

	repository := app.NewRepository(db)
	useCase := app.NewUseCase(repository)
	handler := app.NewHandler(useCase)

	middleware := app.NewMiddleware()

	return &Server{
		handler:    handler,
		middleware: middleware,
	}
}

func (s *Server) Run() *http.Server {
	e := echo.New()
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		he, ok := err.(*echo.HTTPError)
		if ok {
			c.JSON(he.Code, nil)
			return
		}
	}
	customValidationMap := make(map[string]func(fl validator.FieldLevel) bool)
	customValidationMap["date"] = customvalidator.ValidateDateFormat
	customValidator := customvalidator.CustomValidaton(customValidationMap)

	e.Validator = customValidator
	// Allow CORS requests
	e.Use(middleware.CORS())

	NewRouter(e, s.handler, s.middleware)

	appPort := config.GetInt("APP_PORT", port)

	s.http = &http.Server{
		Addr:         fmt.Sprintf(":%d", appPort),
		Handler:      e,
		ReadTimeout:  timeoutServer * time.Second,
		WriteTimeout: timeoutServer * time.Second,
	}

	fmt.Printf("Server running on port %d\n", appPort)

	return s.http
}

func mysqlConnect() (*gorm.DB, error) {
	dbConnection := config.Get("DB_GORM_CONNECTION")
	if dbConnection == "" {
		return nil, errors.New("can't connect to DB")
	}

	dsn := dbConnection

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now()
		},
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return db, err
	}

	fmt.Println("Connection to database established")
	return db, nil
}
