package gorm

import (
	"github.com/gsoultan/uam/core/domain"
	"github.com/gsoultan/uam/core/repository"
	"github.com/jinzhu/gorm"
)

type account struct {
	db *gorm.DB
}

func (this *account) Create(account *domain.Account) error {
	if err := this.db.Create(&account).Error; err != nil {
		return err
	}
	return nil
}

func (this *account) Update(account *domain.Account) error {
	if err := this.db.Updates(&account).Error; err != nil {
		return err
	}
	return nil
}

func (this *account) DeleteByID(id int64) error {
	account, err := this.FindByID(id)
	if err != nil {
		return err
	}
	if err := this.db.Delete(&account).Error; err != nil {
		return err
	}
	return nil
}

func (this *account) FindAll() ([]domain.Account, error) {
	accounts := []domain.Account{}
	if err := this.db.Find(&accounts).Error; err != nil {
		return accounts, err
	}
	return accounts, nil
}

func (this *account) FindByID(id int64) (domain.Account, error) {
	account := domain.Account{}
	if err := this.db.First(account).Error; err != nil {
		return account, err
	}
	return account, nil
}

func NewGormAccount(db *gorm.DB) repository.Account {
	return &account{
		db: db,
	}
}
