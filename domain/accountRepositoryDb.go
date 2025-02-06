package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"

	"bankapp/errs"
	"bankapp/logger"
)

type AccountRepositoryDb struct {
	dbClient *sqlx.DB
}

func (db AccountRepositoryDb) Save(account Account) (*Account, *errs.AppError) {
	accountInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES (?, ?, ?, ? ?)"

	result, err := db.dbClient.Exec(accountInsert, account.CustomerId, account.OpeningDate, account.AccountType, account.Amount, account.Status)

	if err != nil {
		logger.LogError("Error creating account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unepected DB error")
	}

	accountId, err := result.LastInsertId()

	if err != nil {
		logger.LogError("Error while getting last insert id for account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unepected DB error")
	}

	account.AccountId = strconv.FormatInt(accountId, 10)

	return &account, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
