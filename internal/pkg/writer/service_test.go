package writer

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/ramasuryananda/dummy-cv/internal/pkg/exception"
	"github.com/stretchr/testify/assert"
)

func TestAPIResponse(t *testing.T) {
	testCases := []struct {
		name         string
		message      string
		status       bool
		data         interface{}
		expectedJSON string
	}{
		{
			name:         "successful response with data",
			message:      "success",
			status:       true,
			data:         []string{"apple", "banana", "orange"},
			expectedJSON: `{"status":true,"message":"success","data":["apple","banana","orange"]}`,
		},
		{
			name:         "successful response with no data",
			message:      "success",
			status:       true,
			data:         nil,
			expectedJSON: `{"status":true,"message":"success","data":null}`,
		},
		{
			name:         "failed response with error message",
			message:      "error",
			status:       false,
			data:         nil,
			expectedJSON: `{"status":false,"message":"error","data":null}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualResponse := APIResponse(tc.message, tc.status, tc.data)

			// Validate JSON encoding of response
			expectedJSONBytes := []byte(tc.expectedJSON)
			actualJSONBytes, err := json.Marshal(actualResponse)
			if err != nil {
				t.Errorf("failed to marshal actual response to JSON: %v", err)
			}
			assert.Equal(t, expectedJSONBytes, actualJSONBytes, "actual JSON does not match expected JSON")

			// Validate fields of response struct
			assert.Equal(t, tc.status, actualResponse.Status, "actual status does not match expected status")
			assert.Equal(t, tc.message, actualResponse.Message, "actual message does not match expected message")
			assert.Equal(t, tc.data, actualResponse.Data, "actual data does not match expected data")
		})
	}
}

func TestAPIValidationResponse(t *testing.T) {
	testCases := []struct {
		name         string
		message      string
		status       bool
		data         interface{}
		errors       interface{}
		expectedJSON string
	}{
		{
			name:         "successful response with data",
			message:      "success",
			status:       true,
			data:         []string{"apple", "banana", "orange"},
			errors:       nil,
			expectedJSON: `{"status":true,"message":"success","data":["apple","banana","orange"],"errors":null}`,
		},
		{
			name:         "successful response with no data",
			message:      "success",
			status:       true,
			data:         nil,
			errors:       nil,
			expectedJSON: `{"status":true,"message":"success","data":null,"errors":null}`,
		},
		{
			name:         "failed response with error message",
			message:      "error",
			status:       false,
			data:         nil,
			errors:       nil,
			expectedJSON: `{"status":false,"message":"error","data":null,"errors":null}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualResponse := APIValidationResponse(tc.message, tc.status, tc.data, tc.errors)

			// Validate JSON encoding of response
			expectedJSONBytes := []byte(tc.expectedJSON)
			actualJSONBytes, err := json.Marshal(actualResponse)
			if err != nil {
				t.Errorf("failed to marshal actual response to JSON: %v", err)
			}
			assert.Equal(t, expectedJSONBytes, actualJSONBytes, "actual JSON does not match expected JSON")

			// Validate fields of response struct
			assert.Equal(t, tc.status, actualResponse.Status, "actual status does not match expected status")
			assert.Equal(t, tc.message, actualResponse.Message, "actual message does not match expected message")
			assert.Equal(t, tc.data, actualResponse.Data, "actual data does not match expected data")
		})
	}
}

func TestAPIErrorResponse(t *testing.T) {
	type args struct {
		message string
		err     error
	}
	tests := []struct {
		name     string
		args     args
		want     Response
		wantCode int
	}{
		{
			name: "case success - errors format",
			args: args{
				message: "something went wrong",
				err:     errors.New("this is a error message"),
			},
			want: Response{
				Status:  false,
				Message: "something went wrong",
				Data:    nil,
			},
			wantCode: http.StatusInternalServerError,
		},
		{
			name: "case success - exception format",
			args: args{
				message: "something went wrong",
				err:     exception.New("this is a error message"),
			},
			want: Response{
				Status:  false,
				Message: "this is a error message",
				Data:    nil,
			},
			wantCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, gotCode := APIErrorResponse(tt.args.message, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("APIErrorResponse() = %v, want %v", got, tt.want)
				t.Errorf("APIErrorResponse() = %v, want %v", gotCode, tt.wantCode)
			}
		})
	}
}
