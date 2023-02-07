package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	res := Contains([]string{"abc", "def"}, "abc")
	assert.Equal(t, res, true)

	res = Contains([]string{"123", "def"}, "abc")
	assert.Equal(t, res, false)
}
