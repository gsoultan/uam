package app

import "github.com/gsoultan/uam/app"

func Setup() {
	db := Database{}
	if err := app.LoadDatabaseConfig(&db.Config); err != nil {
		return
	}
	d := db.GetDatabaseConnection()
	app.Migrate(d)
}
