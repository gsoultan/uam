package app

import (
	"github.com/caarlos0/env"
	"github.com/gsoultan/uam/config"
	"github.com/joho/godotenv"
	"log"
)

func LoadDatabaseConfig(config *config.Database) error {
	if err := godotenv.Load(); err != nil {
		log.Println("File .env not found, reading configuration from ENV")
		return err
	}

	if err := env.Parse(config); err != nil {
		log.Fatalln("Failed to parse ENV, Err : " + err.Error())
		return err
	}

	return nil
}
