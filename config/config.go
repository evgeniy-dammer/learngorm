package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() (*gorm.DB, error) {
	dsn := "host=localhost user=learngorm password=learngorm dbname=learngorm port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
