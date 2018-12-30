package repository

import "github.com/gsoultan/uam/core/domain"

type Role interface {
	Create(role *domain.Role) error
	Update(role *domain.Role) error
	DeleteByID(id int64) error
	FindAll() ([]domain.Role, error)
	FindByID(id int64) (domain.Role, error)
}
