package repository

import "github.com/clevanilson/cs-trading-platform/account_service/internal/domain/entity"

type AccountRepository interface {
	Save(account entity.Account) error
	Update(account entity.Account) error
	GetByID(id string) (entity.Account, error)
}
