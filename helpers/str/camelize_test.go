package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnderscore(t *testing.T) {
	res := Underscore("UnderScore")
	assert.Equal(t, res, "under_score")

	res = Underscore("UnderScorE")
	assert.Equal(t, res, "under_scor_e")

	res = Underscore("UNDERSCORE")
	assert.Equal(t, res, "underscore")

	res = Underscore("\U0010FFFF")
	assert.Equal(t, res, "\U0010ffff")
}
