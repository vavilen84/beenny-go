package store

import (
	"database/sql"
	"gorm.io/driver/mysql"
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

func GetMockDB(db *sql.DB) (gormDB *gorm.DB) {
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	gormDB = gormDB.Session(&gorm.Session{SkipDefaultTransaction: true})

	return
}
