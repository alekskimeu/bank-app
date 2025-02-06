package dto

import (
	"strings"

	"bankapp/errs"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.AppError {

	if r.Amount < 5000 {
		return errs.NewValidationError("A deposit of at least 5000.00 is required to open an account")
	}

	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		return errs.NewValidationError("Wrong account type")
	}

	return nil

}
