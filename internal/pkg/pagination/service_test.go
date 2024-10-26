package pagination

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestGetOffset(t *testing.T) {
	type args struct {
		page  int
		limit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case get page  more than 1",
			args: args{
				page:  2,
				limit: 10,
			},
			want: 10,
		},
		{
			name: "case get page less than 1",
			args: args{
				page:  0,
				limit: 10,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetOffset(tt.args.page, tt.args.limit)
			if got != tt.want {
				t.Errorf("GetOffset() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPrevPage(t *testing.T) {
	type args struct {
		query string
		page  int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case can get prev page",
			args: args{
				query: "transaction_status=1",
				page:  2,
			},
			want: "page=1&transaction_status=1",
		},
		{
			name: "case can get prev page with no return",
			args: args{
				query: "transaction_status=1",
				page:  1,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new echo context for testing
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/?"+tt.args.query, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if got := GetPrevPage(c, tt.args.page); got != tt.want {
				t.Errorf("GetPrevPage() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestGetNextPage(t *testing.T) {
	type args struct {
		query   string
		page    int
		endrow  int
		maxData int64
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case can get next page",
			args: args{
				query:   "transaction_status=1",
				page:    1,
				endrow:  10,
				maxData: 11,
			},
			want: "page=2&transaction_status=1",
		},
		{
			name: "case can get next page with no return",
			args: args{
				query:   "transaction_status=1",
				page:    1,
				endrow:  10,
				maxData: 10,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new echo context for testing
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/?"+tt.args.query, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if got := GetNextPage(c, tt.args.page, tt.args.endrow, tt.args.maxData); got != tt.want {
				t.Errorf("GetNextPage() = %v, want %v", got, tt.want)
			}
		})
	}
}
