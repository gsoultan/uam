package domain

import (
	"github.com/bwmarrin/snowflake"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type Model struct {
	ID          int64 `gorm:"INDEX;PRIMARY_KEY;UNIQUE;TYPE:BIGINT;NOT NULL;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	CreatedByID *int64 `gorm:"index; null"`
	CreatedBy   *User
	UpdatedByID *int64 `gorm:"index; null"`
	UpdatedBy   *User
	DeletedByID *int64 `gorm:"index; null"`
	DeletedBy   *User
}

func (this *Model) BeforeCreate(scope *gorm.Scope) error {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return err
	}
	if err := scope.SetColumn("ID", node.Generate().Int64()); err != nil {
		return err
	}
	if err := scope.SetColumn("CreatedAt", time.Now()); err != nil {
		return err
	}
	return nil
}

func (this *Model) BeforeSave(scope *gorm.Scope) (err error) {
	if err := scope.SetColumn("UpdatedAt", time.Now()); err != nil {
		return err
	}
	return nil
}
