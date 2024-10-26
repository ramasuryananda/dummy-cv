package translator

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/customvalidator"
)

func TestTranslateError(t *testing.T) {
	type args struct {
		err error
		s   interface{}
	}

	type User struct {
		Username string `query:"username" name:"Username" validate:"required"`
	}

	type Gift struct {
		Tipe string `query:"tipe" name:"Tipe" validate:"required,oneof=buku buka"`
	}

	type ArrayData struct {
		Tipe []string `query:"tipe" name:"Tipe" validate:"required,min=1"`
	}

	type Student struct {
		Nomor    string `query:"nomor" name:"nomor" validate:"required,numeric"`
		Code     string `query:"code" validate:"max=8"`
		ExamDate string `query:"exam_date" name:"Date" validate:"date"`
	}

	req := httptest.NewRequest(http.MethodGet, "/path", nil)
	rec := httptest.NewRecorder()

	// Create a new Echo instance
	e := echo.New()
	// Create a new Echo context using echo.NewContext
	c := e.NewContext(req, rec)

	//add cusstom validation
	customValidationMap := make(map[string]func(fl validator.FieldLevel) bool)
	customValidationMap["date"] = customvalidator.ValidateDateFormat

	customValidator := customvalidator.CustomValidaton(customValidationMap)
	e.Validator = customValidator

	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "Case success mapping required validation error",
			args: args{
				err: c.Validate(User{}),
				s:   User{},
			},
			want: map[string]string{
				"Username": "Username is required",
			},
		},
		{
			name: "Case success mapping oneof validation error",
			args: args{
				err: c.Validate(Gift{Tipe: "mainan"}),
				s:   Gift{},
			},
			want: map[string]string{
				"Tipe": "Tipe is unknown",
			},
		},
		{
			name: "Case success mapping numeric validation error",
			args: args{
				err: c.Validate(Student{
					Nomor: "abcde",
					Code:  "1",
				}),
				s: User{},
			},
			want: map[string]string{
				"Nomor": "Nomor must be numeric",
			},
		},

		{
			name: "Case success mapping max validation error",
			args: args{
				err: c.Validate(Student{
					Nomor: "123456",
					Code:  "123456789",
				}),
				s: User{},
			},
			want: map[string]string{
				"Code": "Code can't be greater than 8",
			},
		},
		{
			name: "Case success mapping date validation error",
			args: args{
				err: c.Validate(Student{
					Nomor: "123456",

					Code:     "12345678",
					ExamDate: "2022-01-01",
				}),
				s: User{},
			},
			want: map[string]string{
				"ExamDate": "ExamDate invalid date format, the date format should be DD/MM/YYYY",
			},
		},
		{
			name: "Case success mapping min validation error",
			args: args{
				err: c.Validate(ArrayData{
					Tipe: []string{},
				}),
				s: User{},
			},
			want: map[string]string{
				"Tipe": "Tipe minimum of 1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TranslateError(tt.args.err, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TranslateError() = %v, want %v", got, tt.want)
			}
		})
	}
}
