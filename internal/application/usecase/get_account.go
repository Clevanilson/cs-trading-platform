package usecase

import "github.com/clevanilson/cs-trading-platform/internal/application/repository"

type GetAccount struct {
	repository repository.AccountRepository
}

func NewGetAccount(repository repository.AccountRepository) *GetAccount {
	return &GetAccount{repository}
}

func (u *GetAccount) Execute(input GetAccountInput) (*GetAccountOutput, error) {
	acount, err := u.repository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}
	return &GetAccountOutput{
		ID:   acount.ID(),
		Name: acount.Name(),
	}, nil
}

type GetAccountInput struct {
	ID string `json:"id"`
}

type GetAccountOutput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
