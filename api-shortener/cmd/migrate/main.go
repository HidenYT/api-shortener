package main

import (
	"api-shortener/storage"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	loadEnv()
	dbSettings := storage.NewDBConnectionSettings()
	db := storage.NewDB(dbSettings)
	migrator := storage.NewMigrator(db)
	migrator.Migrate()
}

const ENV_FILE_NAME = ".env"

func loadEnv() {
	if err := godotenv.Load(ENV_FILE_NAME); err != nil {
		logrus.Fatalf("Couldn't parse load env from %s: %s", ENV_FILE_NAME, err)
	}
}
