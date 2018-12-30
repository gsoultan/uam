package repository

import "github.com/gsoultan/uam/core/domain"

type User interface {
	Create(user *domain.User) error
	Update(user *domain.User) error
	DeleteByID(id int64) error
	SetTokenByID(token string, id int64) error
	SetPasswordByID(password string, id int64) error
	FindAll() ([]domain.User, error)
	FindByAccountID(accountId int64) ([]domain.User, error)
	FindByID(id int64) (domain.User, error)
	FindByUserName(username string) (domain.User, error)
	FindByToken(token string) (domain.User, error)
}
