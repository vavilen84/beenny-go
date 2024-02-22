package models_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/beenny-go/models"
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
	tmpFolder := path.Join(os.Getenv("APP_ROOT"), "tmp", "test_migration")
	p := models.GetMigrationFilePath(name, tmpFolder, now)
	expected := path.Join(tmpFolder, models.GetMigrationFilename(name, now))
	assert.Equal(t, expected, p)
}

func Test_Unit_CreateMigrationFile_ok(t *testing.T) {
	test.LoadEnv()
	now := time.Now()
	name := "create_users_table"
	tmpFolder := path.Join("tmp", "test_migration")
	os.Mkdir(path.Join(os.Getenv("APP_ROOT"), tmpFolder), 0777)
	err := models.CreateMigrationFile(name, tmpFolder, now)
	assert.Nil(t, err)
	p := models.GetMigrationFilePath(name, tmpFolder, now)
	_, err = os.Stat(p)
	assert.False(t, os.IsNotExist(err))
	os.Remove(p)
}

func Test_Unit_getMigration_ok(t *testing.T) {
	test.LoadEnv()
	now := time.Now()
	name := "create_users_table"
	tmpFolder := path.Join("tmp", "test_migration")
	tmpTestAppMigrationsFolder := path.Join(os.Getenv("APP_ROOT"), tmpFolder)
	os.Mkdir(tmpTestAppMigrationsFolder, 0777)
	err := models.CreateMigrationFile(name, tmpFolder, now)
	assert.Nil(t, err)
	p := models.GetMigrationFilePath(name, tmpFolder, now)
	_, err = os.Stat(p)
	assert.False(t, os.IsNotExist(err))

	err = filepath.Walk(tmpTestAppMigrationsFolder, func(path string, info os.FileInfo, err error) error {
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

	os.Remove(p)
}
