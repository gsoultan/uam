package app

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Database struct {
	Config config.Database
}

func (this Database) GetDatabaseConnection() *gorm.DB {
	db, err := gorm.Open(this.Config.Dialect, "host="+this.Config.Host+" port="+this.Config.Port+" user="+this.Config.User+" dbname="+this.Config.Name+" password="+this.Config.Password+" sslmode=disable")
	if err != nil {
		fmt.Print("Error : ", err)
		return nil
	}

	return db
}
