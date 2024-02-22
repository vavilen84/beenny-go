package models_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/beenny-go/models"
	"github.com/vavilen84/beenny-go/test"
	"os"
	"path"
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
	tmpFolder := path.Join(os.Getenv("APP_ROOT"), "tmp", "test_migration")
	err := os.Mkdir(tmpFolder, 0777)
	assert.Nil(t, err)
	err = models.CreateMigrationFile(name, tmpFolder, now)
	assert.Nil(t, err)
	p := models.GetMigrationFilePath(name, tmpFolder, now)
	_, err = os.Stat(p)
	assert.False(t, os.IsNotExist(err))
}

func Test_Unit_getMigration_ok(t *testing.T) {
	//test.LoadEnv()
	//tmpFolderName := path.Join(os.Getenv("APP_ROOT"), "tmp","test_migration")
	//err := os.Mkdir(tmpFolderName, 0777)
	//assert.Nil(t, err)
	//filename :=
}
