package main

import (
	"github.com/joho/godotenv"
	"github.com/vavilen84/nft-project/handlers"
	"github.com/vavilen84/nft-project/store"
	"log"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Can not load .env file")
	}
	store.InitDB()
	handler := handlers.MakeHandler()
	httpServer := handlers.InitHttpServer(handler)
	log.Fatal(httpServer.ListenAndServe())
}
