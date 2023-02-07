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

// SettingDefault ....
type SettingDefault struct {
	AsanaEmailAssignee    string   `json:"asana_email_assignee"`
	AsanaCommentMentionID string   `json:"asana_comment_mention_id"`
	AsanaEmailFollowers   []string `json:"asana_email_followers"`
}

// SettingIncentiveBudget ....
type SettingIncentiveBudget struct {
	OriginAccountNumber      string  `json:"origin_account_number"`
	DestinationAccountNumber string  `json:"destination_account_number"`
	OriginDescription        string  `json:"origin_description"`
	DestinationDescription   string  `json:"destination_description"`
	MinimumBudget            float64 `json:"minimum_budget"`
}

// SettingIncentiveBudgetApproval ....
type SettingIncentiveBudgetApproval struct {
	DebitAccountNumber  string `json:"debit_account_number"`
	CreditAccountNumber string `json:"credit_account_number"`
}

func (m *Setting) ScanRows(row *sql.Row, rows *sql.Rows) error {
	parameters := []interface{}{&m.ID, &m.Name, &m.Value, &m.CreatedAt, &m.UpdatedAt}
	if row != nil {
		return row.Scan(parameters...)
	}
	return rows.Scan(parameters...)
}

var (
	SettingTypeDefault                 = "default_setting"
	SettingTypeIncentiveBudget         = "incentive_budget"
	SettingTypeIncentiveBudgetApproval = "incentive_budget_approval"
	TableName                          = "settings"
	Column                             = []string{`def.id`, `def.name`, `def.value`, `def.created_at`, `def.updated_at`}
	SelectStatement                    = `SELECT ` + strings.Join(Column, ",") + ` FROM ` + TableName + ` def ` + JoinStatement
	SelectCountStatement               = `SELECT COUNT(DISTINCT def.id)  FROM ` + TableName + ` def ` + JoinStatement
	JoinStatement                      = ``
	GroupStatement                     = `GROUP BY def.id`
	OrderBy                            = []string{"def.created_at", "def.updated_at"}
	OrderByString                      = []string{}
	MapOrderBy                         = map[string]string{
		"created_at": "def.created_at",
		"updated_at": "def.updated_at",
	}
)
