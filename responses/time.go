package responses

import (
	"strings"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

type APITime struct {
	time.Time
}

func (t *APITime) UnmarshalJSON(b []byte) error {
	// API return dates in such format: 2021-12-18 19:25:562021-12-18 19:25:56
	trimmedString := strings.Trim(string(b), "\"")
	timezone, err := time.LoadLocation("Europe/Warsaw")
	if err != nil {
		return err
	}
	parsedTime, err := time.ParseInLocation(TimeFormat, trimmedString, timezone)
	if err != nil {
		return err
	}
	*t = APITime{parsedTime.UTC()}
	return nil
}
