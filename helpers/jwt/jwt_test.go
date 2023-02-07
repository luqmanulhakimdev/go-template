package jwt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetToken(t *testing.T) {
	res, err := GetToken("jwttoken")
	assert.IsType(t, res, "jwttoken")
	assert.NoError(t, err)
}
