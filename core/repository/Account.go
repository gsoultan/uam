package repository

import "github.com/gsoultan/uam/core/domain"

type Account interface {
	Create(account *domain.Account) error
	Update(account *domain.Account) error
	DeleteByID(id int64) error
	FindAll() ([]domain.Account, error)
	FindByID(id int64) (domain.Account, error)
}
