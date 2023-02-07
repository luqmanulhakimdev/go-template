package time

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInString(t *testing.T) {
	res, _ := InString("2022-04-10", "2006-01-02", "Asia/Jakarta")
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err == nil {
		res = res.In(loc)
	}

	assert.Equal(t, res, time.Date(2022, time.April, 10, 7, 0, 0, 0, loc))

	res, err = InString("2006-01-02 00:00:)0", "2006-01-02", "Asia/Jakarta")
	assert.Error(t, err)
}

func TestIn(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Jakarta")

	res, _ := In(time.Date(2022, time.April, 10, 0, 0, 0, 0, time.UTC), "Asia/Jakarta")
	assert.Equal(t, res, time.Date(2022, time.April, 10, 7, 0, 0, 0, loc))
}

func TestConvert(t *testing.T) {
	res := Convert("2022-04-10", "2006-01-02", "2006 01 02")
	assert.Equal(t, res, "2022 04 10")

	res = Convert("2022-04-10 00:00:00", "2006-01-02", "2006 01 02")
	assert.Equal(t, res, "")
}

func TestConvertWithTimezone(t *testing.T) {
	res := ConvertWithTimezone("2022-04-10", "2006-01-02", "2006 01 02", "Asia/Jakarta")
	assert.Equal(t, res, "2022 04 10")

	res = ConvertWithTimezone("2022-04-10 00:00:00", "2006-01-02", "2006 01 02", "Asia/Jakarta")
	assert.Equal(t, res, "")

	res = ConvertWithTimezone("2022-04-10", "2006-01-02", "2006 01 02", "Asia")
	assert.Equal(t, res, "")
}
