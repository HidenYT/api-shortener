package storage_test

import (
	"api-shortener/shortreq"
	"api-shortener/storage"
	"os"
	"testing"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var testShortenedAPIDAO *storage.ShortenedAPIDAO
var testShorteningRuleDAO *storage.ShorteningRuleDAO
var testOutgoingRequestHeaderDAO *storage.OutgoingRequestHeaderDAO
var testOutgoingRequestParamDAO *storage.OutgoingRequestParamDAO
var testOutgoingRequestConfigDAO *storage.OutgoingRequestConfigDAO

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env.test")
	if err != nil {
		logrus.Fatal("cannot load .env.test file:", err)
	}

	settings := storage.DBCreationSettings{}
	err = env.Parse(&settings)
	if err != nil {
		logrus.Fatal("cannot parse DBCreationSettings from env:", err)
	}

	db, err := storage.CreateDB(&settings)
	if err != nil {
		logrus.Fatal("cannot connect to db:", err)
	}
	migrator := storage.NewMigrator(db)
	migrator.Migrate()
	validator := shortreq.NewValidate()
	testShortenedAPIDAO = storage.NewShortenedAPIDAO(db, validator)
	testShorteningRuleDAO = storage.NewShorteningRuleDAO(db, validator)
	testOutgoingRequestHeaderDAO = storage.NewOutgoingRequestHeaderDAO(db, validator)
	testOutgoingRequestParamDAO = storage.NewOutgoingRequestParamDAO(db, validator)
	testOutgoingRequestConfigDAO = storage.NewOutgoingRequestConfigDAO(db, validator)

	code := m.Run()
	os.Clearenv()
	os.Exit(code)
}
