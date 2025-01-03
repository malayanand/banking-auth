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
	GenerateAndSaveRefreshTokenToStore(AuthToken) (string, *errs.AppError)
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

func (d AuthRepositoryDb) GenerateAndSaveRefreshTokenToStore(authToken AuthToken) (string, *errs.AppError) {
	var appErr *errs.AppError
	var refreshToken string

	if refreshToken, appErr = authToken.newRefreshToken(); appErr != nil {
		return "", appErr
	}

	// store it in store
	sqlInsert := "insert into refresh_token_store (refresh_token) values (?)"
	_, err := d.client.Exec(sqlInsert, refreshToken)
	if err != nil {
		logger.Error("Unexpected database error:" + err.Error())
		return "", errs.NewUnexpectedError("Unexpected database error")
	}

	return refreshToken, nil
}

func NewAuthRepository(client *sqlx.DB) AuthRepositoryDb {
	return AuthRepositoryDb{client: client}
}
