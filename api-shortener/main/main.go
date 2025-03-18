package main

import (
	"api-shortener/http"
	"api-shortener/shortreq"
	"api-shortener/storage"
)

func main() {
	LoadEnv()

	validator := shortreq.NewValidate()

	dbSettings := storage.NewDBConnectionSettings()
	db := storage.NewDB(dbSettings)
	migrator := storage.NewMigrator(db)
	migrator.Migrate()

	apiDAO := storage.NewShortenedAPIDAO(db, validator)
	configDAO := storage.NewOutgoingRequestConfigDAO(db, validator)
	headerDAO := storage.NewOutgoingRequestHeaderDAO(db, validator)
	paramDAO := storage.NewOutgoingRequestParamDAO(db, validator)
	ruleDAO := storage.NewShorteningRuleDAO(db, validator)

	apiService := http.NewAPIService(apiDAO)
	configService := http.NewRequestConfigService(configDAO)
	headerService := http.NewRequestHeaderService(headerDAO)
	paramService := http.NewRequestParamService(paramDAO)
	ruleService := http.NewShorteningRuleService(ruleDAO)

	apiClientSettings := http.NewOutgoingRequestClientSettings()
	apiClient := http.NewOutgoingRequestClient(apiClientSettings)
	limiterSettings := http.NewLoopLimiterSettings()
	limiter := http.NewLoopLimiter(limiterSettings)
	shorteningService := http.NewResponseShorteningService(configDAO, headerDAO, paramDAO, apiClient, limiter)

	server := http.NewHTTPServer(
		apiDAO, shorteningService, apiService, configService, headerService, paramService, ruleService,
	)

	server.Run()
}
