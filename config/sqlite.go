package config

import (
	"os"

	"github.com/admin-srf/go_crud/schemas"
	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

func InitSqlite() (*gorm.DB, error) {
	dbPath := "./data.db"

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}
	db, err := gorm.Open(sqlite.Open("./data.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&schemas.User{})
	if err != nil {
		return nil, err
	}
	return db, nil

}
