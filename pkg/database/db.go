package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))
	pg := postgres.Open(dsn)
	db, err := gorm.Open(pg, &gorm.Config{})
	if err == nil {
		if err := db.AutoMigrate(&Movie{}, &Review{}); err != nil {
			return nil, err
		}
	}
	return db, err
}

func Must(db *gorm.DB, err error) *gorm.DB {
	if err != nil {
		log.Fatal(err)
	}
	return db
}
