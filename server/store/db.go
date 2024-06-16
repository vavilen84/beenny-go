package store

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
)

var (
	db *gorm.DB
)

func InitDB() {
	db = initDb()
}

func InitTestDB() {
	db = initTestDb()
}

func InitSQLServerConnection() {
	db = initSqlServerConnection()
}

func GetDB() *gorm.DB {
	return db
}

func initDb() *gorm.DB {
	return processInitDb(os.Getenv("DB_SQL_DSN"))
}

func initTestDb() *gorm.DB {
	return processInitDb(os.Getenv("TEST_DB_SQL_DSN"))
}

func initSqlServerConnection() *gorm.DB {
	return processInitDb(os.Getenv("SQL_SERVER_DSN"))
}

func processInitDb(DbDsn string) (db *gorm.DB) {
	db, err := gorm.Open(mysql.Open(DbDsn), &gorm.Config{})
	if err != nil {
		panic("failed to database: " + err.Error())
	}
	db.Debug()
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
	return
}
