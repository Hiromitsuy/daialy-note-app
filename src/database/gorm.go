package database

import (
	"fmt"
	"os"
	"strconv"

	"github.com/lunasky-hy/dialy-note-app/src/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres() *gorm.DB {
	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")
	DB_PORTNUM, _ := strconv.Atoi(DB_PORT)

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Tokyo", 
		DB_HOST, DB_USER, DB_NAME, DB_PORTNUM)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if (err != nil) {
		fmt.Println(err.Error())
	}

	db.AutoMigrate(&model.Question{})
	return db
}