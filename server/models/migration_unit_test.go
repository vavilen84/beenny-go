package models_test

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/beenny-go/helpers"
	"github.com/vavilen84/beenny-go/mocks"
	"github.com/vavilen84/beenny-go/models"
	"github.com/vavilen84/beenny-go/store"
	"github.com/vavilen84/beenny-go/test"
	"os"
	"path"
	"path/filepath"
	"testing"
	"time"
)

func Test_Unit_GetMigrationFilename_ok(t *testing.T) {
	now := time.Now()
	name := "create_users_table"
	filename := models.GetMigrationFilename(name, now)
	expected := fmt.Sprintf("%d_%s.up.sql", now.Unix(), name)
	assert.Equal(t, expected, filename)
}

func Test_Unit_GetMigrationFilePath_ok(t *testing.T) {
	now := time.Now()
	name := "create_users_table"
	tmpFolder := getTestMigrationsFolder()
	p := models.GetMigrationFilePath(name, tmpFolder, now)
	expected := path.Join(tmpFolder, models.GetMigrationFilename(name, now))
	assert.Equal(t, expected, p)
}

func getTestMigrationsFolder() string {
	return path.Join(os.Getenv("APP_ROOT"), "tmp", "test_migration")
}

func createTestMigrationFile(t *testing.T, name string, tt time.Time) string {
	tmpFolder := getTestMigrationsFolder()
	os.Mkdir(tmpFolder, 0777)
	err := models.CreateMigrationFile(name, tmpFolder, tt)
	assert.Nil(t, err)
	p := models.GetMigrationFilePath(name, tmpFolder, tt)
	_, err = os.Stat(p)
	assert.False(t, os.IsNotExist(err))
	return models.GetMigrationFilePath(name, tmpFolder, tt)
}

func Test_Unit_CreateMigrationFile_ok(t *testing.T) {
	test.LoadEnv()
	now := time.Now()
	name := "create_users_table"
	p := createTestMigrationFile(t, name, now)
	os.Remove(p)
}

func Test_Unit_GetMigration_ok(t *testing.T) {
	test.LoadEnv()
	now := time.Now()
	name := "create_users_table"
	p := createTestMigrationFile(t, name, now)
	err := filepath.Walk(getTestMigrationsFolder(), func(path string, info os.FileInfo, err error) error {
		assert.Nil(t, err)
		if !info.IsDir() {
			err, m := models.GetMigration(info)
			assert.Nil(t, err)
			assert.Equal(t, m.Version, now.Unix())
			assert.Equal(t, m.Filename, models.GetMigrationFilename(name, now))
			assert.NotEmpty(t, m.CreatedAt)
		}
		return nil
	})
	assert.Nil(t, err)
	os.Remove(p)
}

func Test_Unit_GetMigrations_ok(t *testing.T) {
	test.LoadEnv()
	now := time.Now()
	name := "create_users_table"
	migration1 := createTestMigrationFile(t, name, now)

	future := time.Now().Add(1 * time.Hour)
	name2 := "create_jwt_info_table"
	migration2 := createTestMigrationFile(t, name2, future)

	err, keys, list := models.GetMigrations(getTestMigrationsFolder())
	assert.Nil(t, err)
	assert.Equal(t, 2, len(keys))
	assert.Equal(t, int(now.Unix()), keys[0])
	assert.Equal(t, int(future.Unix()), keys[1])

	assert.Equal(t, list[now.Unix()].Version, now.Unix())
	assert.Equal(t, list[now.Unix()].Filename, models.GetMigrationFilename(name, now))
	assert.NotEmpty(t, list[now.Unix()].CreatedAt)

	assert.Equal(t, list[future.Unix()].Version, future.Unix())
	assert.Equal(t, list[future.Unix()].Filename, models.GetMigrationFilename(name2, future))
	assert.NotEmpty(t, list[future.Unix()].CreatedAt)

	os.Remove(migration1)
	os.Remove(migration2)
}

func Test_Unit_CreateMigrationsTableIfNotExists_ok(t *testing.T) {
	customMatcher := mocks.CustomMatcher{}
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	gormDB := store.GetMockDB(db)

	sql := "CREATE TABLE IF NOT EXISTS migrations"
	sqlMock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(0, 1))

	err = models.CreateMigrationsTableIfNotExists(gormDB)
	assert.Nil(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_Unit_InsertMigration_ok(t *testing.T) {
	customMatcher := mocks.CustomMatcher{}
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	gormDB := store.GetMockDB(db)

	sql := "INSERT INTO `migrations`"
	sqlMock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(1, 1))

	now := time.Now()
	m := models.Migration{
		CreatedAt: time.Now().Unix(),
		Filename:  "create_users_table",
		Version:   now.Unix(),
	}
	err = models.InsertMigration(gormDB, &m)
	assert.Nil(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_Unit_PerformMigrateTx_ok(t *testing.T) {
	test.LoadEnv()
	customMatcher := mocks.CustomMatcher{}
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	gormDB := store.GetMockDB(db)

	sqlMock.ExpectBegin()

	sql := "INSERT INTO `migrations`"
	sqlMock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(1, 1))

	sql = "CREATE TABLE `users`"
	sqlMock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(1, 1))

	sqlMock.ExpectCommit()

	m := models.Migration{
		CreatedAt: int64(1708526378),
		Filename:  "1708526378_add_users_table.up.sql",
		Version:   int64(1708526378),
	}
	err = models.PerformMigrateTx(gormDB, m, helpers.GetFixturesFolder())
	assert.Nil(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_Unit_Apply_ok(t *testing.T) {
	test.LoadEnv()
	customMatcher := mocks.CustomMatcher{}
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	gormDB := store.GetMockDB(db)

	rows := sqlmock.NewRows([]string{"id"}) // no rows
	expectedSQL := "SELECT * FROM `migrations`"
	sqlMock.ExpectQuery(expectedSQL).WillReturnRows(rows)

	sqlMock.ExpectBegin()

	sql := "INSERT INTO `migrations`"
	sqlMock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(1, 1))

	sql = "CREATE TABLE `users`"
	sqlMock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(1, 1))

	sqlMock.ExpectCommit()

	m := models.Migration{
		CreatedAt: int64(1708526378),
		Filename:  "1708526378_add_users_table.up.sql",
		Version:   int64(1708526378),
	}
	k := 1708526378
	list := make(map[int64]models.Migration)
	list[int64(1708526378)] = m
	err = models.Apply(gormDB, k, list, helpers.GetFixturesFolder())
	assert.Nil(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_Unit_MigrateUp_ok(t *testing.T) {

}
