package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatCurrency(t *testing.T) {
	res := FormatCurrency(float64(10000), "IDR", ".", ",", 0)
	assert.Equal(t, res, "10.000")
}
