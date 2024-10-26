package formatter

import (
	"fmt"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/ramasuryananda/dummy-cv/internal/constant"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func CurrencyFormat(amount float64) string {
	if amount == float64(int64(amount)) {
		return humanize.FormatInteger("#.###,", int(amount))
	}

	return humanize.FormatFloat("#.###,##", amount)
}

func TimeToUnixTime(t *time.Time) int64 {
	return t.Unix()
}

func FormattedDateToString(dateString string, layout string, format string) string {
	// Parse the input date string into a time.Time value
	// layout example time.RFC3339Nano
	date, err := time.Parse(layout, dateString)
	if err != nil {
		fmt.Println("Error parsing date formattedDateToString() function: ", err)
		return dateString
	}

	// Format the time.Time value
	// example "Monday, 2 January 2006 15:04"
	formattedDate := date.Format(format)

	return formattedDate
}

func FormattingDate(dateString string) (time.Time, error) {
	// Parse the input date string using the input format
	date, err := time.Parse(constant.DateFormatDDMMYYY, dateString)
	if err != nil {
		date, _ := time.Parse("2006-01-01", "0000-00-00")
		return date, err
	}

	return date, nil
}

func CapitalizeString(s string) string {
	s = strings.ReplaceAll(s, "_", " ")
	words := strings.Fields(s)
	caser := cases.Title(language.English)
	for i, word := range words {
		words[i] = caser.String(word)
	}
	return strings.Join(words, " ")
}
