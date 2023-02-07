package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateProcessID(t *testing.T) {
	res := GenerateProcessID()
	assert.IsType(t, res, "abcdefgh")
}

func TestSetPaginationParameter(t *testing.T) {
	offset, limit, page, orderBy, sort := SetPaginationParameter(1, 10, "date", "asc", []string{"date"}, []string{})
	assert.Equal(t, offset, 0)
	assert.Equal(t, limit, 10)
	assert.Equal(t, page, 1)
	assert.Equal(t, orderBy, "date")
	assert.Equal(t, sort, "asc")

	_, limit, page, orderBy, sort = SetPaginationParameter(0, 0, "name", "", []string{"name"}, []string{"name"})
	assert.Equal(t, limit, 10)
	assert.Equal(t, page, 1)
	assert.Equal(t, orderBy, "LOWER(name)")
	assert.Equal(t, sort, "asc")

	_, _, _, orderBy, _ = SetPaginationParameter(0, 0, "", "", []string{}, []string{})
	assert.Equal(t, orderBy, "def.updated_at")
}

func TestcheckWhiteList(t *testing.T) {
	res := checkWhiteList("name", []string{"name"})
	assert.Equal(t, res, "name")

	res = checkWhiteList("name", []string{})
	assert.Equal(t, res, "def.updated_at")
}
