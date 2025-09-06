package api_dao_test

import (
	"os"
	"testing"

	api_dao "github.com/HidenYT/api-shortener/internal/storage/dao/api"
	"github.com/HidenYT/api-shortener/internal/storage/migration"
	"github.com/HidenYT/api-shortener/internal/validation"

	"github.com/HidenYT/api-shortener/internal/storage"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var testShortenedAPIDAO *api_dao.ShortenedAPIDAO
var testShorteningRuleDAO *api_dao.ShorteningRuleDAO
var testOutgoingRequestHeaderDAO *api_dao.OutgoingRequestHeaderDAO
var testOutgoingRequestParamDAO *api_dao.OutgoingRequestParamDAO
var testOutgoingRequestConfigDAO *api_dao.OutgoingRequestConfigDAO

func TestMain(m *testing.M) {
	err := godotenv.Load("../../../.env.test")
	if err != nil {
		logrus.Fatal("cannot load .env.test file:", err)
	}

	settings := storage.DBConnectionSettings{}
	err = env.Parse(&settings)
	if err != nil {
		logrus.Fatal("cannot parse DBCreationSettings from env:", err)
	}

	db := storage.NewDB(&settings)
	migrator := migration.NewAPIDBMigrator(db)
	migrator.Migrate()
	validator := validation.NewValidate()
	testShortenedAPIDAO = api_dao.NewShortenedAPIDAO(db, validator)
	testShorteningRuleDAO = api_dao.NewShorteningRuleDAO(db, validator)
	testOutgoingRequestHeaderDAO = api_dao.NewOutgoingRequestHeaderDAO(db, validator)
	testOutgoingRequestParamDAO = api_dao.NewOutgoingRequestParamDAO(db, validator)
	testOutgoingRequestConfigDAO = api_dao.NewOutgoingRequestConfigDAO(db, validator)

	code := m.Run()
	os.Clearenv()
	os.Exit(code)
}
