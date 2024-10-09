package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabaseConnection() *gorm.DB {
	dsnURL := "host=db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsnURL), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
