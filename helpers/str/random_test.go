package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandStringBytesMaskImprSrc(t *testing.T) {
	res := RandStringBytesMaskImprSrc(4)
	assert.IsType(t, res, "abcdefgh")
}

func TestRandAlphanumericString(t *testing.T) {
	res := RandAlphanumericString(4)
	assert.IsType(t, res, "abc123")
}

func TestRandLowerAlphanumericString(t *testing.T) {
	res := RandLowerAlphanumericString(4)
	assert.IsType(t, res, "abc123")
}

func TestStringWithCharset(t *testing.T) {
	res := StringWithCharset(4, "abcdefgh")
	assert.IsType(t, res, "abcdefgh")
}

func TestRandomNumericString(t *testing.T) {
	res := RandomNumericString(4)
	assert.IsType(t, res, "12345")
}
