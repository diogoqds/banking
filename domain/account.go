package domain

import (
	"github.com/diogoqds/banking/dto"
	"github.com/diogoqds/banking/errs"
)

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
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		AccountId: a.AccountId,
	}
}
