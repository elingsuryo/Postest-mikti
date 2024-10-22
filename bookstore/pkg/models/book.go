package models

import (
	"bookstore/pkg/config" // Pastikan jalur ini benar

	"gorm.io/gorm"
)

// Model Buku
type Book struct {
	gorm.Model
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

// Fungsi untuk melakukan migrasi model Book di database
func Migrate() {
	if config.DB == nil {
		panic("Database connection is not initialized!")
	}
	config.DB.AutoMigrate(&Book{}) // Lakukan migrasi di sini
}