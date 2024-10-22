package main

import (
	"bookstore/pkg/config"
	"bookstore/pkg/models" // Pastikan ini diimpor untuk memanggil fungsi Migrate
	"bookstore/pkg/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// Inisialisasi Echo
	e := echo.New()

	// Load konfigurasi database
	config.Connect() // Panggil ini untuk inisialisasi database

	// Lakukan migrasi model
	models.Migrate() // Panggil ini untuk migrasi model

	// Setup routes
	routes.BookStoreRoutes(e)

	// Start server pada port 8080
	e.Logger.Fatal(e.Start(":8080"))
}