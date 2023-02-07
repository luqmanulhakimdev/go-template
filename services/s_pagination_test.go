package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetPaginationParameter(t *testing.T) {
	offset, limit, page, orderBy, sort := SetPaginationParameter(1, 10, "created_at", "desc", []string{"created_at"}, []string{"created_at"})
	assert.Equal(t, offset, 0)
	assert.Equal(t, limit, 10)
	assert.Equal(t, page, 1)
	assert.Equal(t, orderBy, "LOWER(created_at)")
	assert.Equal(t, sort, "desc")

	offset, limit, page, orderBy, sort = SetPaginationParameter(0, 0, "id", "ascc", []string{}, []string{"created_at"})
	assert.Equal(t, offset, 0)
	assert.Equal(t, limit, 10)
	assert.Equal(t, page, 1)
	assert.Equal(t, orderBy, "def.updated_at")
	assert.Equal(t, sort, "asc")

	_, _, _, orderBy, _ = SetPaginationParameter(1, 10, "updated_at", "desc", []string{"created_at"}, []string{"created_at"})
	assert.Equal(t, orderBy, "LOWER(created_at)")
}

func TestSetPaginationResponse(t *testing.T) {
	res := SetPaginationResponse(1, 10, 10)
	assert.Equal(t, res, Pagination{CurrentPage: 1, LastPage: 1, Total: 10, PerPage: 10})

	res = SetPaginationResponse(1, 10, 0)
	assert.Equal(t, res, Pagination{CurrentPage: 1, LastPage: 0, Total: 0, PerPage: 10})
}
