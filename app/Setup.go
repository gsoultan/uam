package app

func Setup() {
	db := Database{}
	if err := LoadDatabaseConfig(&db.Config); err != nil {
		return
	}
	d := db.GetDatabaseConnection()
	Migrate(d)
}
