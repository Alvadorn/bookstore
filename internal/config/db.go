package config

import (
	"github.com/alvadorn/bookstore/pkg/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// ConnectDB creates a connect to database
func ConnectDB() (*gorm.DB, error) {
	return gorm.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=bookstore sslmode=disable")
}

// AutoMigrate function
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Book{})
}
