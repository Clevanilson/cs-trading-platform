package infrarepository

import (
	"github.com/clevanilson/cs-trading-platform/internal/domain/entity"
)

type accountMemoryRepository struct {
	data map[string]entity.Account
}

func NewAccountMemoryRepository() *accountMemoryRepository {
	return &accountMemoryRepository{
		data: make(map[string]entity.Account),
	}
}

func (r *accountMemoryRepository) Save(account entity.Account) error {
	r.data[account.ID()] = account
	return nil
}

func (r *accountMemoryRepository) Update(account entity.Account) error {
	r.data[account.ID()] = account
	return nil
}

func (r *accountMemoryRepository) GetByID(id string) (entity.Account, error) {
	account, ok := r.data[id]
	if !ok {
		return nil, nil
	}
	return account, nil
}
