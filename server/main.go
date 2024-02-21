package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/vavilen84/beenny-go/constants"
	"github.com/vavilen84/beenny-go/handlers"
	"github.com/vavilen84/beenny-go/helpers"
	"github.com/vavilen84/beenny-go/models"
	"github.com/vavilen84/beenny-go/store"
	"log"
	"os"
	"path/filepath"
	"time"
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
	&cobra.Command{
		Use:  "migration-create",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			store.InitDB()
			now := time.Now()
			nowUnix := now.Unix()

			file := filepath.Join(constants.MigrationsFolder, fmt.Sprintf("%d_%s.up.sql", nowUnix, args[0]))
			_, err := os.Create(file)
			if err != nil {
				helpers.LogError(err)
			}
		},
	},
	&cobra.Command{
		Use:  "migrate-up",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			store.InitDB()
			db := store.GetDB()
			err := models.CreateMigrationsTableIfNotExists(db)
			if err != nil {
				log.Println(err)
			}
			err = models.MigrateUp(db)
			if err != nil {
				log.Println(err)
			}
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
