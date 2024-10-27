package general

import (
	"encoding/json"
	"time"
)

type YMDDate time.Time

const customLayout = "02-01-2006"

func (date YMDDate) MarshalJSON() ([]byte, error) {
	timeData := time.Time(date)

	if timeData.Equal(time.Time{}) {
		return json.Marshal("")
	}

	return json.Marshal(timeData.Format(customLayout))
}
