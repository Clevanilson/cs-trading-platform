package usecase

import (
	"github.com/clevanilson/cs-trading-platform/account_service/internal/application/repository"
	pkgerror "github.com/clevanilson/cs-trading-platform/devpack/pkg/error"
)

type GetAccount interface {
	Execute(input GetAccountInput) (*GetAccountOutput, error)
}

type GetAccountInput struct {
	ID string `json:"id"`
}

type GetAccountOutput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type getAccount struct {
	repository repository.AccountRepository
}

func NewGetAccount(repository repository.AccountRepository) *getAccount {
	return &getAccount{repository}
}

func (u *getAccount) Execute(input GetAccountInput) (*GetAccountOutput, error) {
	account, err := u.repository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, pkgerror.NewNotFound("Account")
	}
	return &GetAccountOutput{
		ID:   account.ID(),
		Name: account.Name(),
	}, nil
}
