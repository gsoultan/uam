package main

import (
	"github.com/gsoultan/uam/app"
)

func main() {
	db := app.Database{}
	if err := app.LoadConfiguration(&db.Config); err != nil {
		return
	}
	d := db.GetDatabaseConnection()
	app.Migrate(d)
}
