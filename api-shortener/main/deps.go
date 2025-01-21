package main

import (
	"api-shortener/restapi"
	"api-shortener/security"
	"api-shortener/shortreq"
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

const ENV_FILE_NAME = ".env"

func NewHTTPServer(
	lc fx.Lifecycle,
	shorteningService shortreq.IResponseShorteningService,
	apiRepo shortreq.IShortenedAPIDAO,
	restService restapi.IRESTService,
) *http.Server {
	ginServer := gin.Default()
	ginServer.Use(security.APITokenChecker())
	shortreq.AttachAPIShorteningGroup(ginServer, shorteningService, apiRepo)
	restapi.AttachRESTAPIGroup(ginServer, restService)
	return appendGinToLifecycle(lc, ginServer)
}

func appendGinToLifecycle(lc fx.Lifecycle, r *gin.Engine) *http.Server {
	srv := &http.Server{Addr: ":8080", Handler: r.Handler()}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			fmt.Println("Starting HTTP server at", srv.Addr)
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}

func LoadEnv() {
	if err := godotenv.Load(ENV_FILE_NAME); err != nil {
		panic(err)
	}
}
