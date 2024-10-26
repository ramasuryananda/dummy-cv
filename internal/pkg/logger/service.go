package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime"

	"github.com/google/uuid"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/config"
	"github.com/sirupsen/logrus"
)

type key string

const (
	envDirName       = "LOG_DIR"
	envEnableLogging = "LOG_ENABLE"
	envLoggingMode   = "LOG_MODE"

	keyAppName       = "app"
	keyRequestID key = "request_id"
)

var (
	appName = ""
)

// Init is used to initilize log.
func Init(name string) {
	dir := config.Get(envDirName)
	enableLogging := config.Get(envEnableLogging)
	loggingMode := config.GetWIthDefault(envLoggingMode, "file")

	path := dir + "-" + name + ".log"
	appName = name // set app name

	logrus.SetLevel(logrus.TraceLevel)

	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if enableLogging == "enable" {

		if loggingMode == "stdout" {
			// Output to stdout instead of the default stderr
			// Can be any io.Writer, see below for File example
			logrus.SetOutput(os.Stdout)
		}

		if loggingMode == "file" {
			// Output to file instead
			logFile, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				log.Fatal(err)
			}
			logrus.SetOutput(logFile)
		}
	}
}

// InitLogCtx initializes log contex that used for easy logging trace
// by using unique request id for identifier
func InitLogCtx(ctx context.Context) context.Context {
	if GetRequestID(ctx) == "" {
		return SetRequestID(ctx, uuid.NewString())
	}

	return ctx
}

// Info is used to logging level info
func Info(ctx context.Context, metadata interface{}, err error, message string) {
	logrus.WithFields(logrus.Fields(toMap(ctx, metadata, err))).Info(message)
}

// Warning is used to logging level warning
func Warning(ctx context.Context, metadata interface{}, err error, message string) {
	logrus.WithFields(logrus.Fields(toMap(ctx, metadata, err))).Warning(message)
}

// Error is used to logging level error
func Error(ctx context.Context, metadata interface{}, err error, message string) {
	logrus.WithFields(logrus.Fields(toMap(ctx, metadata, err))).Error(message)
}

// Fatal is used to logging level fatal and program will exit
func Fatal(ctx context.Context, metadata interface{}, err error, message string) {
	logrus.WithFields(logrus.Fields(toMap(ctx, metadata, err))).Fatal(message)
}

// Trace is used to logging level trace
func Trace(ctx context.Context, metadata interface{}, err error, message string) {
	logrus.WithFields(logrus.Fields(toMap(ctx, metadata, err))).Trace(message)
}

// toMap is used to convert value to map
func toMap(ctx context.Context, val interface{}, err error) map[string]interface{} {
	result := make(map[string]interface{})
	result["request_id"] = GetRequestID(ctx)
	result["app"] = appName

	_, dirFile, lineNum, _ := runtime.Caller(2)
	result["file"] = fmt.Sprintf("%s:%d", dirFile, lineNum)

	if err != nil {
		result["error"] = err.Error()
	}

	if val == nil {
		return result
	}

	dataType := reflect.ValueOf(val).Kind()
	switch dataType {
	case reflect.Struct:
		valStruct := struct {
			Metadata interface{} `json:"metadata"`
		}{val}
		data, _ := json.Marshal(valStruct)
		json.Unmarshal(data, &result)
	default:
		result["metadata"] = val
	}

	return result
}

// GetRequestID gets request id from context
func GetRequestID(ctx context.Context) string {
	requestID, ok := ctx.Value(keyRequestID).(string)
	if !ok {
		return ""
	}

	return requestID
}

// SetRequestID sets request id to context
func SetRequestID(ctx context.Context, randomKey string) context.Context {
	return context.WithValue(ctx, keyRequestID, randomKey)
}
