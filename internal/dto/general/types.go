package general

import (
	"encoding/json"
	"time"
)

type YMDDate time.Time

const customLayout = "02-01-2006"

func (date YMDDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(date).Format(customLayout))
}
