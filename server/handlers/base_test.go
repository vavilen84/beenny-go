package handlers_test

import (
	"github.com/joho/godotenv"
	"github.com/vavilen84/beenny-go/handlers"
	"github.com/vavilen84/beenny-go/helpers"
	"github.com/vavilen84/beenny-go/models"
	"github.com/vavilen84/beenny-go/store"
	"gorm.io/gorm"
	"log"
	"net/http/httptest"
	"os"
)

var (
	testAppInited = false
)

func loadEnv() (err error) {
	// we run test from current folder
	if err := godotenv.Load(".env"); err != nil {
		// we run test from root folder
		err = godotenv.Load("../.env")
		if err != nil {
			panic("can not load env file")
		}
	}
	return err
}

func afterEachTest(conn *gorm.DB) {
	sess, _ := conn.DB()
	sess.Close()
}

func beforeEachTest() {
	if !testAppInited {
		err := loadEnv()
		if err != nil {
			panic(err)
		}
		testAppInited = true
	}
	err := clearTestDb()
	if err != nil {
		panic(err)
	}
	migrate()
}

func clearTestDb() error {
	store.InitSQLServerConnection()
	db := store.GetDB()
	db.Exec("DROP DATABASE beenny_test")
	err := db.Exec("CREATE DATABASE beenny_test").Error
	if err != nil {
		panic(err)
	}
	return err
}

func getAppPath() string {
	appPath := os.Getenv("APP_ROOT_PATH")
	if appPath == "" {
		panic("APP_ROOT_PATH env var is empty")
	}
	return appPath
}

func migrate() {
	store.InitDB()
	db := store.GetDB()
	err := models.CreateMigrationsTableIfNotExists(db)
	if err != nil {
		log.Println(err)
	}
	err = models.MigrateUp(db, helpers.GetMigrationsFolder())
	if err != nil {
		log.Println(err)
	}
}

func makeTestServer() *httptest.Server {
	handler := handlers.MakeHandler()
	ts := httptest.NewServer(handler)
	return ts
}
