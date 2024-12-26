package domain

import "database/sql"

type Login struct {
	Username   string         `db:"username"`
	CustomerId sql.NullString `db:"customer_id"`
	Accounts   sql.NullString `db:"accounts"`
	Role       string         `db:"role"`
}
