package gorm

import (
	"github.com/gsoultan/uam/core/domain"
	"github.com/gsoultan/uam/core/repository"
	"github.com/jinzhu/gorm"
)

type role struct {
	db *gorm.DB
}

func (this *role) Create(role *domain.Role) error {
	if err := this.db.Create(&role).Error; err != nil {
		return err
	}
	return nil
}

func (this *role) Update(role *domain.Role) error {
	if err := this.db.Update(&role).Error; err != nil {
		return err
	}
	return nil
}

func (this *role) DeleteByID(id int64) error {
	role, err := this.FindByID(id)
	if err != nil {
		return err
	}

	if err = this.db.Delete(&role).Error; err != nil {
		return err
	}

	return nil
}

func (this *role) FindAll() ([]domain.Role, error) {
	roles := []domain.Role{}
	if err := this.db.Find(&roles).Error; err != nil {
		return []domain.Role{}, err
	}
	return roles, nil
}

func (this *role) FindByID(id int64) (domain.Role, error) {
	role := domain.Role{}
	if err := this.db.First(&role, id).Error; err != nil {
		return domain.Role{}, err
	}
	return role, nil
}

func NewRoleGorm(db *gorm.DB) repository.Role {
	return &role{
		db: db,
	}
}
