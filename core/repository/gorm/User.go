package gorm

import (
	"github.com/gsoultan/uam/core/domain"
	"github.com/gsoultan/uam/core/repository"
	"github.com/jinzhu/gorm"
)

type user struct {
	db *gorm.DB
}

func (this *user) FindByToken(token string) (domain.User, error) {
	user := domain.User{}
	if err := this.db.Where(&domain.User{Token: token}).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (this *user) FindByUserName(username string) (domain.User, error) {
	user := domain.User{}
	if err := this.db.Where(&domain.User{UserName: username}).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (this *user) SetPasswordByID(password string, id int64) error {
	user := domain.User{}
	user.ID = id
	user.Password = password

	if err := this.db.Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

func (this *user) SetTokenByID(token string, id int64) error {
	user := domain.User{}
	user.ID = id

	if err := this.db.Model(&user).Update("token", token).Error; err != nil {
		return err
	}
	return nil
}

func (this *user) Create(user *domain.User) error {
	if err := this.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (this *user) Update(user *domain.User) error {
	if err := this.db.Model(user).Updates(domain.User{UserName: user.UserName, Email: user.Email, Mobile: user.Mobile}).Error; err != nil {
		return err
	}
	return nil
}

func (this *user) DeleteByID(id int64) error {
	user, err := this.FindByID(id)
	if err != nil {
		return err
	}
	if err := this.db.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func (this *user) FindAll() ([]domain.User, error) {
	users := []domain.User{}
	if err := this.db.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func (this *user) FindByAccountID(accountId int64) ([]domain.User, error) {
	users := []domain.User{}
	if err := this.db.Where(&domain.User{AccountID: accountId}).Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func (this *user) FindByID(id int64) (domain.User, error) {
	user := domain.User{}
	if err := this.db.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func NewUserGorm(db *gorm.DB) repository.User {
	return &user{
		db: db,
	}
}
