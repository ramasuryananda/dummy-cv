package formatter

import (
	"testing"
	"time"
)

func TestCurrencyFormat(t *testing.T) {
	type args struct {
		amount float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case format currency 1000 into 1.000",
			args: args{
				amount: 1000,
			},
			want: "1.000",
		},
		{
			name: "case format currency 1000.01 into 1.000,01",
			args: args{
				amount: 1000.01,
			},
			want: "1.000,01",
		},
		{
			name: "case format currency 100 into 100",
			args: args{
				amount: 100,
			},
			want: "100",
		},
		{
			name: "case format currency 1234567 into 1.234.567",
			args: args{
				amount: 1234567,
			},
			want: "1.234.567",
		},
		{
			name: "case format currency 1234567.89 into 1.234.567,89",
			args: args{
				amount: 1234567.89,
			},
			want: "1.234.567,89",
		},
		{
			name: "case format currency 100.009 into 100,01",
			args: args{
				amount: 100.009,
			},
			want: "100,01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CurrencyFormat(tt.args.amount); got != tt.want {
				t.Errorf("CurrencyFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapitalizeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case success 1",
			args: args{
				s: "nasi_goreng",
			},
			want: "Nasi Goreng",
		},
		{
			name: "case success 2",
			args: args{
				s: "nasi goreng",
			},
			want: "Nasi Goreng",
		},
		{
			name: "case success 3",
			args: args{
				s: "golang",
			},
			want: "Golang",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CapitalizeString(tt.args.s); got != tt.want {
				t.Errorf("CapitalizeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeToUnixTime(t *testing.T) {
	type args struct {
		t *time.Time
	}

	time := time.Now()

	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "case success",
			args: args{
				t: &time,
			},
			want: time.Unix(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeToUnixTime(tt.args.t); got != tt.want {
				t.Errorf("TimeToUnixTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormattedDateToString(t *testing.T) {
	type args struct {
		dateString string
		layout     string
		format     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case success",
			args: args{
				dateString: "2006-01-02T15:04:05",
				layout:     "2006-01-02T15:04:05",
				format:     "Monday, 2 January 2006 15:04",
			},
			want: "Monday, 2 January 2006 15:04",
		},
		{
			name: "case error",
			args: args{
				dateString: "2006-01-02T15:04:05",
				layout:     "?",
				format:     "Monday, 2 January 2006 15:04",
			},
			want: "2006-01-02T15:04:05",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormattedDateToString(tt.args.dateString, tt.args.layout, tt.args.format); got != tt.want {
				t.Errorf("FormattedDateToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormattingDate(t *testing.T) {
	type args struct {
		dateString string
	}

	date, _ := time.Parse("2006-01-02", "2006-01-02")
	date2, _ := time.Parse("2006-01-02", "2023-12-31")
	failDate, _ := time.Parse("2006-01-02", "0000-00-00")

	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name: "case success",
			args: args{
				dateString: "02/01/2006",
			},
			want:    date,
			wantErr: false,
		},
		{
			name: "case success - 2",
			args: args{
				dateString: "31/12/2023",
			},
			want:    date2,
			wantErr: false,
		},
		{
			name: "case error",
			args: args{
				dateString: "02-01-2006",
			},
			want:    failDate,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FormattingDate(tt.args.dateString)
			if (err != nil) != tt.wantErr {
				t.Errorf("FormattingDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FormattingDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
