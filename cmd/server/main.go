package main

import (
	"github.com/HidenYT/api-shortener/internal/http"
	crudapi_v1 "github.com/HidenYT/api-shortener/internal/http/crudapi/v1"
	crudapi_v2 "github.com/HidenYT/api-shortener/internal/http/crudapi/v2"
	http_shortener "github.com/HidenYT/api-shortener/internal/http/shorten"
	shortener "github.com/HidenYT/api-shortener/internal/response-shortener"
	"github.com/HidenYT/api-shortener/internal/storage"
	api_dao "github.com/HidenYT/api-shortener/internal/storage/dao/api"
	"github.com/HidenYT/api-shortener/internal/validation"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	loadEnv()

	validator := validation.NewValidate()

	dbSettings := storage.NewDBConnectionSettings()
	db := storage.NewDB(dbSettings)

	apiDAO := api_dao.NewShortenedAPIDAO(db, validator)
	configDAO := api_dao.NewOutgoingRequestConfigDAO(db, validator)
	headerDAO := api_dao.NewOutgoingRequestHeaderDAO(db, validator)
	paramDAO := api_dao.NewOutgoingRequestParamDAO(db, validator)
	ruleDAO := api_dao.NewShorteningRuleDAO(db, validator)

	apiService := crudapi_v1.NewAPIService(apiDAO)
	configService := crudapi_v1.NewRequestConfigService(configDAO)
	headerService := crudapi_v1.NewRequestHeaderService(headerDAO)
	paramService := crudapi_v1.NewRequestParamService(paramDAO)
	ruleService := crudapi_v1.NewShorteningRuleService(ruleDAO)

	apiClientSettings := shortener.NewOutgoingRequestClientSettings()
	apiClient := shortener.NewOutgoingRequestClient(apiClientSettings)
	responseShortener := shortener.NewResponseShortener(apiClient)

	limiterSettings := http_shortener.NewLoopLimiterSettings()
	limiter := http_shortener.NewLoopLimiter(limiterSettings)
	shorteningService := http_shortener.NewResponseShorteningService(configDAO, headerDAO, paramDAO, responseShortener, limiter)

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
