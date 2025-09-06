package main

import (
	"api-shortener/http"
	crudapi_v1 "api-shortener/http/crudapi/v1"
	crudapi_v2 "api-shortener/http/crudapi/v2"
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

	apiDAO := shortreq.NewShortenedAPIDAO(db, validator)
	configDAO := shortreq.NewOutgoingRequestConfigDAO(db, validator)
	headerDAO := shortreq.NewOutgoingRequestHeaderDAO(db, validator)
	paramDAO := shortreq.NewOutgoingRequestParamDAO(db, validator)
	ruleDAO := shortreq.NewShorteningRuleDAO(db, validator)

	apiService := crudapi_v1.NewAPIService(apiDAO)
	configService := crudapi_v1.NewRequestConfigService(configDAO)
	headerService := crudapi_v1.NewRequestHeaderService(headerDAO)
	paramService := crudapi_v1.NewRequestParamService(paramDAO)
	ruleService := crudapi_v1.NewShorteningRuleService(ruleDAO)

	apiClientSettings := shortener.NewOutgoingRequestClientSettings()
	apiClient := shortener.NewOutgoingRequestClient(apiClientSettings)
	responseShortener := shortener.NewResponseShortener(apiClient)

	limiterSettings := crudapi_v1.NewLoopLimiterSettings()
	limiter := crudapi_v1.NewLoopLimiter(limiterSettings)
	shorteningService := crudapi_v1.NewResponseShorteningService(configDAO, headerDAO, paramDAO, responseShortener, limiter)

	apiDTOService := crudapi_v2.NewAPIDTOService(apiDAO, configDAO, ruleDAO, headerDAO, paramDAO)

	server := http.NewHTTPServer(
		apiDAO, shorteningService, apiService, configService, headerService, paramService, ruleService, apiDTOService,
	)

	server.Run()
}

const ENV_FILE_NAME = ".env"

func loadEnv() {
	if err := godotenv.Load(ENV_FILE_NAME); err != nil {
		logrus.Fatalf("Couldn't parse load env from %s: %s", ENV_FILE_NAME, err)
	}
}
