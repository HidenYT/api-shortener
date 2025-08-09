package shortreq_test

import (
	"api-shortener/shortreq"
	"api-shortener/storage"
	"os"
	"testing"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var testShortenedAPIDAO *shortreq.ShortenedAPIDAO
var testShorteningRuleDAO *shortreq.ShorteningRuleDAO
var testOutgoingRequestHeaderDAO *shortreq.OutgoingRequestHeaderDAO
var testOutgoingRequestParamDAO *shortreq.OutgoingRequestParamDAO
var testOutgoingRequestConfigDAO *shortreq.OutgoingRequestConfigDAO

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env.test")
	if err != nil {
		logrus.Fatal("cannot load .env.test file:", err)
	}

	settings := storage.DBConnectionSettings{}
	err = env.Parse(&settings)
	if err != nil {
		logrus.Fatal("cannot parse DBCreationSettings from env:", err)
	}

	db := storage.NewDB(&settings)
	migrator := storage.NewMigrator(db)
	migrator.Migrate()
	validator := shortreq.NewValidate()
	testShortenedAPIDAO = shortreq.NewShortenedAPIDAO(db, validator)
	testShorteningRuleDAO = shortreq.NewShorteningRuleDAO(db, validator)
	testOutgoingRequestHeaderDAO = shortreq.NewOutgoingRequestHeaderDAO(db, validator)
	testOutgoingRequestParamDAO = shortreq.NewOutgoingRequestParamDAO(db, validator)
	testOutgoingRequestConfigDAO = shortreq.NewOutgoingRequestConfigDAO(db, validator)

	code := m.Run()
	os.Clearenv()
	os.Exit(code)
}
