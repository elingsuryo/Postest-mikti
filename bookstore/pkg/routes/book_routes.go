package routes

import (
	"bookstore/pkg/controllers" // Pastikan jalur impor sesuai dengan struktur proyek

	"github.com/labstack/echo/v4"
)

// Fungsi untuk mendefinisikan semua route bookstore
func BookStoreRoutes(e *echo.Echo) {
	e.GET("/books", controllers.GetBooks)          // Mendapatkan semua buku
	e.GET("/books/:id", controllers.GetBookByID)   // Mendapatkan buku berdasarkan ID
	e.POST("/books", controllers.CreateBook)       // Membuat buku baru
	e.PUT("/books/:id", controllers.UpdateBook)    // Mengupdate buku berdasarkan ID
	e.DELETE("/books/:id", controllers.DeleteBook) // Menghapus buku berdasarkan ID
}