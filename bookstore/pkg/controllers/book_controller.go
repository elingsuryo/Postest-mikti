package controllers

import (
	"bookstore/pkg/config"
	"bookstore/pkg/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Create Book
func CreateBook(c echo.Context) error {
	var book models.Book
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	config.DB.Create(&book)
	return c.JSON(http.StatusOK, book)
}

// Get All Books
func GetBooks(c echo.Context) error {
	var books []models.Book
	config.DB.Find(&books)
	return c.JSON(http.StatusOK, books)
}

// Get Book by ID
func GetBookByID(c echo.Context) error {
	id := c.Param("bookid")
	var book models.Book
	config.DB.First(&book, id)
	if book.ID == 0 {
		return c.JSON(http.StatusNotFound, "Book not found")
	}
	return c.JSON(http.StatusOK, book)
}

// Update Book
func UpdateBook(c echo.Context) error {
	id := c.Param("bookid")
	var book models.Book
	config.DB.First(&book, id)
	if book.ID == 0 {
		return c.JSON(http.StatusNotFound, "Book not found")
	}
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	config.DB.Save(&book)
	return c.JSON(http.StatusOK, book)
}

// Delete Book
func DeleteBook(c echo.Context) error {
	id := c.Param("bookid")
	var book models.Book
	config.DB.Delete(&book, id)
	return c.JSON(http.StatusOK, "Book deleted")
}