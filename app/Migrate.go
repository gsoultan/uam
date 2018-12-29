package app

import (
	"fmt"
	"github.com/gsoultan/uam/domain"
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&domain.User{})
	fmt.Println("Create Table UserAccount")

	db.AutoMigrate(&domain.Page{})
	fmt.Println("Create Table Page")

	db.AutoMigrate(&domain.Account{})
	fmt.Println("Create Table Account")

	db.AutoMigrate(&domain.Company{})
	fmt.Println("Create Table Company")

	db.AutoMigrate(&domain.Role{})
	fmt.Println("Create Table Role")

	db.AutoMigrate(&domain.RoleInModule{})
	fmt.Println("Create Table Role in Module")

	db.AutoMigrate(&domain.Menu{})
	fmt.Println("Create Table Menu")
}
