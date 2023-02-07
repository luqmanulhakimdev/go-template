package str

import (
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
)

func TestDateStringToDate(t *testing.T) {
	res := DateStringToDate("2022-04-10", "2006-01-02")
	assert.Equal(t, res, time.Date(2022, time.April, 10, 0, 0, 0, 0, time.UTC))
}
