package store

import (
	"database/sql"
	"github.com/vavilen84/beenny-go/env"
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

func GetDB() *gorm.DB {
	return db
}

func initDb() *gorm.DB {
	DbDsn := os.Getenv("DB_SQL_DSN")
	return processInitDb(DbDsn)
}

func initTestDb() *gorm.DB {
	DbDsn := os.Getenv("TEST_DB_SQL_DSN")
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

func processInitDb(DbDsn string) (db *gorm.DB) {
	db, err := gorm.Open(mysql.Open(DbDsn), &gorm.Config{})
	if err != nil {
		panic("failed to database: " + err.Error())
	}
	if env.IsDevelopmentEnv() {
		db.Debug()
	}
	sqlDB, err := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
	return
}
