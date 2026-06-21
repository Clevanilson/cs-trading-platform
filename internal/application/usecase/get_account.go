package usecase

import (
	"github.com/clevanilson/cs-trading-platform/internal/application/repository"
	"github.com/clevanilson/cs-trading-platform/pkg/errorc"
)

type GetAccount struct {
	repository repository.AccountRepository
}

func NewGetAccount(repository repository.AccountRepository) *GetAccount {
	return &GetAccount{repository}
}

func (u *GetAccount) Execute(input GetAccountInput) (*GetAccountOutput, error) {
	account, err := u.repository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, errorc.NewNotFound("Account")
	}
	return &GetAccountOutput{
		ID:   account.ID(),
		Name: account.Name(),
	}, nil
}

type GetAccountInput struct {
	ID string `json:"id"`
}

type GetAccountOutput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
