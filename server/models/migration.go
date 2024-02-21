package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/vavilen84/beenny-go/constants"
	"github.com/vavilen84/beenny-go/helpers"
	"github.com/vavilen84/beenny-go/validation"
	"gorm.io/gorm"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Migration struct {
	Id        uint32 `json:"id"`
	Version   int64  `json:"version"`
	Filename  string `json:"filename"`
	CreatedAt int64  `json:"created_at"`
}

func (Migration) GetTableName() string {
	return "migrations"
}

func (Migration) GetValidationRules() interface{} {
	return validation.ScenarioRules{
		constants.ScenarioCreate: validation.FieldRules{
			"Version":   "required",
			"Filename":  "required",
			"CreatedAt": "required",
		},
	}
}

func (Migration) GetValidator() interface{} {
	return validator.New()
}

func getMigration(info os.FileInfo) (err error, m Migration) {
	filename := info.Name()
	splitted := strings.Split(info.Name(), "_")
	version, err := strconv.Atoi(splitted[0])
	if err != nil {
		helpers.LogError(err)
		return
	}

	m = Migration{
		CreatedAt: time.Now().Unix(),
		Filename:  filename,
		Version:   int64(version),
	}
	return
}

func getMigrations() (err error, keys []int, list map[int64]Migration) {
	list = make(map[int64]Migration)
	keys = make([]int, 0)
	folder := path.Join(os.Getenv("APP_ROOT"), constants.MigrationsFolder)
	err = filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			helpers.LogError(err)
		}
		if !info.IsDir() {
			err, migration := getMigration(info)
			if err != nil {
				helpers.LogError(err)
				return err
			}
			list[migration.Version] = migration
			keys = append(keys, int(migration.Version))
		}
		return nil
	})
	if err != nil {
		helpers.LogError(err)
		return
	}

	sort.Ints(keys)
	return
}

func MigrateUp(db *gorm.DB) error {
	err, keys, list := getMigrations()
	for _, k := range keys {
		err = apply(db, k, list)
		if err != nil {
			helpers.LogError(err)
			return err
		}
	}
	return nil
}

func performMigrateTx(db *gorm.DB, m Migration) error {
	tx := db.Begin()
	if tx.Error != nil {
		helpers.LogError(tx.Error)
		return tx.Error
	}
	err := InsertMigration(tx, &m)
	if err != nil {
		helpers.LogError(tx.Error)
		return tx.Error
	}

	filename := path.Join(os.Getenv("APP_ROOT"), constants.MigrationsFolder, m.Filename)
	content, readErr := os.ReadFile(filename)
	if readErr != nil {
		tx.Rollback()
		helpers.LogError(readErr)
		return readErr
	}
	sqlQuery := string(content)
	err = tx.Exec(sqlQuery).Error
	if err != nil {
		tx.Rollback()
		helpers.LogError(err)
		return err
	}

	if err := tx.Commit().Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func apply(db *gorm.DB, k int, list map[int64]Migration) error {
	m := list[int64(k)]
	mm := Migration{}
	err := db.Where("version = ?", m.Version).First(&mm).Error
	if err == gorm.ErrRecordNotFound {
		validationErr := validation.ValidateByScenario(constants.ScenarioCreate, m)
		if validationErr != nil {
			helpers.LogError(validationErr)
			return validationErr
		}
		err = performMigrateTx(db, m)
		if err != nil {
			helpers.LogError(err)
			return err
		}
	} else if err != nil {
		return err
	}
	return nil
}

func InsertMigration(db *gorm.DB, m *Migration) (err error) {
	err = validation.ValidateByScenario(constants.ScenarioCreate, *m)
	if err != nil {
		helpers.LogError(err)
		return
	}
	err = db.Create(m).Error
	if err != nil {
		helpers.LogError(err)
	}
	return
}

func CreateMigrationsTableIfNotExists(db *gorm.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS migrations
		(
   		    id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
			version BIGINT UNSIGNED NOT NULL,
			filename varchar(255) NOT NULL,
			created_at BIGINT UNSIGNED NOT NULL
		) ENGINE=InnoDB CHARSET=utf8;
	`
	err := db.Exec(query).Error
	return err
}
