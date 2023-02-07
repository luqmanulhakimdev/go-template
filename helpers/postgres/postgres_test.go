package postgres

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceSQL(t *testing.T) {
	res := ReplaceSQL(`INSERT INTO table (name, description) VALUES %s`, "(?, ?)", 1)
	assert.Equal(t, res, "INSERT INTO table (name, description) VALUES ($1, $2)")
}

func TestSubstitutePlaceholder(t *testing.T) {
	res := SubstitutePlaceholder("SELECT * FROM table WHERE id = ?", 1)
	assert.Equal(t, res, "SELECT * FROM table WHERE id = $1")
}

func TestBuildStrParams(t *testing.T) {
	res := BuildStrParams(4, ",")
	assert.Equal(t, res, "?,?,?,?")
}

func TestAppendStrArgs(t *testing.T) {
	var res []interface{}
	AppendStrArgs([]string{"John", "Doe"}, &res)
	assert.Equal(t, res, []interface{}{"John", "Doe"})
}

func TestAppendIntArgs(t *testing.T) {
	var res []interface{}
	AppendIntArgs([]int{123, 456}, &res)
	assert.Equal(t, res, []interface{}{123, 456})
}
