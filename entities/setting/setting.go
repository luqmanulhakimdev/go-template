package setting

import (
	"database/sql"
	"strings"
)

// Setting ....
type Setting struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Value     string `json:"value"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (m *Setting) ScanRows(row *sql.Row, rows *sql.Rows) error {
	parameters := []interface{}{&m.ID, &m.Name, &m.Value, &m.CreatedAt, &m.UpdatedAt}
	if row != nil {
		return row.Scan(parameters...)
	}
	return rows.Scan(parameters...)
}

var (
	TableName            = "settings"
	Column               = []string{`def.id`, `def.name`, `def.value`, `def.created_at`, `def.updated_at`}
	SelectStatement      = `SELECT ` + strings.Join(Column, ",") + ` FROM ` + TableName + ` def ` + JoinStatement
	SelectCountStatement = `SELECT COUNT(DISTINCT def.id)  FROM ` + TableName + ` def ` + JoinStatement
	JoinStatement        = ``
	GroupStatement       = `GROUP BY def.id`
	OrderBy              = []string{"def.created_at", "def.updated_at"}
	OrderByString        = []string{}
	MapOrderBy           = map[string]string{
		"created_at": "def.created_at",
		"updated_at": "def.updated_at",
	}
)
