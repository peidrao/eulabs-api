package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "user:password@tcp(localhost:3306)/eulabs"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})

	if err != nil {
		panic("Can't connect to database")
	}

	return db
}
