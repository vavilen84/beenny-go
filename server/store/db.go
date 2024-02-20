package store

import (
	"gorm.io/gorm"
	"os"
)

var (
	db *gorm.DB
)

func InitDB() {
	db = initDb()
}

func GetDB() *gorm.DB {
	return db
}

func initDb() *gorm.DB {
	DbDsn := os.Getenv("DB_SQL_DSN")
	return processInitDb(DbDsn)
}
