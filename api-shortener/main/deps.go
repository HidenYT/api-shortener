package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

const ENV_FILE_NAME = ".env"

func LoadEnv() {
	if err := godotenv.Load(ENV_FILE_NAME); err != nil {
		logrus.Fatalf("Couldn't parse load env from %s: %s", ENV_FILE_NAME, err)
	}
}
