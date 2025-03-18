package main

import (
	"github.com/joho/godotenv"
)

const ENV_FILE_NAME = ".env"

func LoadEnv() {
	if err := godotenv.Load(ENV_FILE_NAME); err != nil {
		panic(err)
	}
}
