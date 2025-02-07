package domain

import (
	"bankapp/dto"
	"bankapp/errs"
)

const dbTSLayout = "2006-01-02 15:04:05"

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	SaveTransaction(Transaction) (*Transaction, *errs.AppError)
	FindBy(accountId string) (*Account, *errs.AppError)
}

func (account Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountId: account.AccountId}
}

func (a Account) CanWithdraw(amount float64) bool {
	return a.Amount > amount
}
