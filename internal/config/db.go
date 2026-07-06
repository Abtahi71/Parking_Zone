package config

import "fmt"

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb(cfg *Config) *gorm.DB {
	dsn := cfg.Dsn
	db,err:=gorm.Open(postgres.Open(dsn),&gorm.Config{
		TranslateError:true,
	})

	if err != nil {
		panic("Failed to connect to database")
	} else {
		fmt.Println("Connected to database")
	}

	return db
}