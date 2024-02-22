package test

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		err := godotenv.Load("../.env")
		if err != nil {
			log.Fatal("Can not load .env file")
		}
	}
}
