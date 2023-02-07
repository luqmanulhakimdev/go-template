package time

import (
	"time"
)

const (
	JAKARTA_TIME_LOCATION = "Asia/Jakarta"
	DATE_FORMAT           = "2006-01-02"
)

// In ...
func InString(t, format, locName string) (time.Time, error) {
	date, err := time.Parse(format, t)
	if err != nil {
		return date, err
	}

	loc, err := time.LoadLocation(locName)
	if err == nil {
		date = date.In(loc)
	}
	return date, err
}

// In ...
func In(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}

// Convert ...
func Convert(t, fromFormat, toFormat string) string {
	timeConvert, err := time.Parse(fromFormat, t)
	if err != nil {
		return ""
	}

	return timeConvert.Format(toFormat)
}

// ConvertWithTimezone ...
func ConvertWithTimezone(t, fromFormat, toFormat, timeZone string) string {
	timeConvert, err := time.Parse(fromFormat, t)
	if err != nil {
		return ""
	}

	localTime, err := In(timeConvert, timeZone)
	if err != nil {
		return ""
	}

	return localTime.Format(toFormat)
}
