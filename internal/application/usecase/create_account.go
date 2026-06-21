package usecase

import (
	"github.com/clevanilson/cs-trading-platform/internal/application/repository"
	"github.com/clevanilson/cs-trading-platform/internal/domain/entity"
)

type CreateAccount struct {
	repository repository.AccountRepository
}

func NewCreateAccount(repository repository.AccountRepository) *CreateAccount {
	return &CreateAccount{repository}
}

func (u *CreateAccount) Execute(input CreateAccountInput) (*CreateAccountOutput, error) {
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

type CreateAccountInput struct {
	Name string `json:"name"`
}

type CreateAccountOutput struct {
	ID string `json:"id"`
}
