package main

import (
	"api-shortener/http"
	shortener "api-shortener/response-shortener"
	"api-shortener/shortreq"
	"api-shortener/storage"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	loadEnv()

	validator := shortreq.NewValidate()

	dbSettings := storage.NewDBConnectionSettings()
	db := storage.NewDB(dbSettings)
	migrator := storage.NewMigrator(db)
	migrator.Migrate()

	apiDAO := shortreq.NewShortenedAPIDAO(db, validator)
	configDAO := shortreq.NewOutgoingRequestConfigDAO(db, validator)
	headerDAO := shortreq.NewOutgoingRequestHeaderDAO(db, validator)
	paramDAO := shortreq.NewOutgoingRequestParamDAO(db, validator)
	ruleDAO := shortreq.NewShorteningRuleDAO(db, validator)

	apiService := http.NewAPIService(apiDAO)
	configService := http.NewRequestConfigService(configDAO)
	headerService := http.NewRequestHeaderService(headerDAO)
	paramService := http.NewRequestParamService(paramDAO)
	ruleService := http.NewShorteningRuleService(ruleDAO)

	apiClientSettings := shortener.NewOutgoingRequestClientSettings()
	apiClient := shortener.NewOutgoingRequestClient(apiClientSettings)
	responseShortener := shortener.NewResponseShortener(apiClient)

	limiterSettings := http.NewLoopLimiterSettings()
	limiter := http.NewLoopLimiter(limiterSettings)
	shorteningService := http.NewResponseShorteningService(configDAO, headerDAO, paramDAO, responseShortener, limiter)

	server := http.NewHTTPServer(
		apiDAO, shorteningService, apiService, configService, headerService, paramService, ruleService,
	)

	server.Run()
}

const ENV_FILE_NAME = ".env"

func loadEnv() {
	if err := godotenv.Load(ENV_FILE_NAME); err != nil {
		logrus.Fatalf("Couldn't parse load env from %s: %s", ENV_FILE_NAME, err)
	}
}
