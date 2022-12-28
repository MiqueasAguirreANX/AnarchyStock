package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

var DB Database

const (
	PAGE_SIZE = 10
)

func InitDB() {
	var err error
	Production := os.Getenv("PRODUCTION")

	if Production == "true" {
		DbHost := os.Getenv("DB_HOST")
		DbUser := os.Getenv("DB_USER")
		DbPassword := os.Getenv("DB_PASSWORD")
		DbName := os.Getenv("DB_NAME")
		DbPort := os.Getenv("DB_PORT")
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			DbHost,
			DbUser,
			DbPassword,
			DbName,
			DbPort,
		)
		DB.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Println("failed to connect db")
			os.Exit(1)
		}
		fmt.Println(DB.DB)
	} else {
		DB.DB, err = gorm.Open(sqlite.Open("db.sqlite3"), &gorm.Config{})
		if err != nil {
			fmt.Println("failed to connect db")
			os.Exit(1)
		}
		fmt.Println(DB.DB)
	}
}

func Paginator(db_model interface{}, page uint) ([]map[string]interface{}, int64, int64) {
	if page < 1 {
		page = 1
	}
	var results []map[string]interface{}
	var total, totalPages int64
	DB.DB.Model(db_model).Count(&total)
	totalPages = total / PAGE_SIZE
	lastPageCount := total % PAGE_SIZE
	if lastPageCount != 0 {
		totalPages += 1
	}
	if page > uint(totalPages) {
		page = uint(totalPages)
	}
	DB.DB.Model(db_model).Limit(PAGE_SIZE).Offset(int(page-1) * PAGE_SIZE).Find(&results)
	return results, total, totalPages
}
