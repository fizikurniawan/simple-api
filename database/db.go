package database

import (
	"fmt"
	"os"
	"simple-api/api/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var err error

func Init() {
	var (
		DB_HOST    = os.Getenv("DB_HOST")
		DB_USER    = os.Getenv("DB_USER")
		DB_NAME    = os.Getenv("DB_NAME")
		DB_PORT    = os.Getenv("DB_PORT")
		DB_SSLMODE = os.Getenv("DB_SSLMODE")
		DB_TZ      = os.Getenv("DB_TZ")
	)

	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		DB_HOST, DB_USER, DB_NAME, DB_PORT, DB_SSLMODE, DB_TZ,
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to connect database")
	}

	// Migration
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Food{})
}

func DbManager() *gorm.DB {
	return db
}
