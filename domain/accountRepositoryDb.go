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

func (db AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	// starting the DB transaction block
	tx, err := db.dbClient.Begin()

	if err != nil {
		logger.LogError("Error while starting a new transaction for bank account transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected DB error")
	}

	// insert bank account transaction
	result, _ := tx.Exec(`INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) VALUES (?, ?, ?, ?)`, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	// update account balance
	if t.IsWithdrawal() {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount - ? WHERE account_id = ?`, t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount + ? WHERE account_id = ?`, t.Amount, t.AccountId)
	}

	// in case of an error, rollback -> changes from both tbles will be reverted
	if err != nil {
		tx.Rollback()
		logger.LogError("Error while saving transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected DB error")
	}

	// commit the transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.LogError("Error while commiting transaction for bank account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected DB error")
	}

	// get the last transaction id from the transactions table
	transactionId, err := result.LastInsertId()

	if err != nil {
		logger.LogError("Error while getting last transaction id: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected DB error")
	}

	// get the latest account info from the accounts table
	account, appErr := db.FindBy(t.AccountId)

	if appErr != nil {
		return nil, appErr
	}

	t.TransactionId = strconv.FormatInt(transactionId, 10)

	// update the transaction struct with the latest balance
	t.Amount = account.Amount

	return &t, nil
}

func (d AccountRepositoryDb) FindBy(accountId string) (*Account, *errs.AppError) {
	sqlGetAccount := "SELECT account_id, customer_id, opening_date, account_type, amount from accounts where account_id = ?"
	var account Account
	err := d.dbClient.Get(&account, sqlGetAccount, accountId)
	if err != nil {
		logger.LogError("Error while fetching account information: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &account, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
