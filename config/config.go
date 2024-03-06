package config

import "gorm.io/gorm"

var (
	db *gorm.DB
)

func InitConfig() error {
	var err error

	db, err = InitSqlite()

	if err != nil {
		return err
	}

	return nil
}

func GetDb() *gorm.DB {
	return db
}
