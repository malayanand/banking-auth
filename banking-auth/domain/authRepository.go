package domain

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/malayanand/banking/errs"
	"github.com/malayanand/banking/logger"
)

type AuthRepositoryDb struct {
	client *sqlx.DB
}

type AuthRepository interface {
	FindBy(string, string) (*Login, *errs.AppError)
}

func (d AuthRepositoryDb) FindBy(username string, password string) (*Login, *errs.AppError) {
	var login Login
	sqlVerify := `SELECT username, u.customer_id, role, group_concat(a.account_id) as account_numbers FROM users u LEFT JOIN accounts a ON a.customer_id = u.customer_id WHERE username=? AND password=? GROUP BY a.customer_id`

	err := d.client.Get(&login, sqlVerify, username, password)
	if err != nil {
		if err == sql.ErrNoRows {
			errs.NewAuthenticationError("invalid credentials")
		} else {
			logger.Error("Error while verifying login credentials from database: " + err.Error())
			errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &login, nil
}

func NewAuthRepository(client *sqlx.DB) AuthRepositoryDb {
	return AuthRepositoryDb{client: client}
}
