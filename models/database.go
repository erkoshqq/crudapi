package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=online_library password=admin sslmode=disable")
	if err != nil {
		panic("Failed to connect to the database")
	}
	db.AutoMigrate(&Book{})

	return db
}
