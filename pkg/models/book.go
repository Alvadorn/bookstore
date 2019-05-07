package models

import "github.com/jinzhu/gorm"

// Book model
type Book struct {
	gorm.Model
	Name string
	ISBN string
}
