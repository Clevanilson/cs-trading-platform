package usecase

import (
	"github.com/clevanilson/cs-trading-platform/account_service/internal/application/repository"
	"github.com/clevanilson/cs-trading-platform/account_service/internal/domain/entity"
)

type CreateAccount interface {
	Execute(input CreateAccountInput) (*CreateAccountOutput, error)
}

type CreateAccountInput struct {
	Name string `json:"name"`
}

type CreateAccountOutput struct {
	ID string `json:"id"`
}

type createAccount struct {
	repository repository.AccountRepository
}

func NewCreateAccount(repository repository.AccountRepository) *createAccount {
	return &createAccount{repository}
}

func (u *createAccount) Execute(input CreateAccountInput) (*CreateAccountOutput, error) {
	account, err := entity.NewAccount(entity.AccountBuilder{
		Name: input.Name,
	})
	if err != nil {
		return nil, err
	}
	err = u.repository.Save(account)
	if err != nil {
		return nil, err
	}
	return &CreateAccountOutput{ID: account.ID()}, nil
}
