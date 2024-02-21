package main

import (
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/vavilen84/beenny-go/handlers"
	"github.com/vavilen84/beenny-go/store"
	"log"
)

var AppCommands = []*cobra.Command{
	&cobra.Command{
		Use: "run-server",
		Run: func(cmd *cobra.Command, args []string) {
			store.InitDB()
			handler := handlers.MakeHandler()
			httpServer := handlers.InitHttpServer(handler)
			log.Fatal(httpServer.ListenAndServe())
		},
	},
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var rootCmd = &cobra.Command{}
	for _, command := range AppCommands {
		rootCmd.AddCommand(command)
	}
	rootCmd.Execute()
}
