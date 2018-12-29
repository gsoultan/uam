package domain

import "time"

type Account struct {
	Model
	Name             string    `gorm:"UNIQUE_INDEX UNIQUE" json:"name"`
	Email            string    `gorm:"UNIQUE_INDEX UNIQUE" json:"email"`
	CompanyID        int64     `gorm:"INDEX" json:"-"`
	Company          Company   `json:"company"`
	RegistrationDate time.Time `json:"registration_date"`
	Users            []User    `json:"users"`
	isActive         bool      `json:"is_active"`
}
