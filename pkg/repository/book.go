package repository

import (
	"serengeti.app/go-rest-template/pkg/db"
	"serengeti.app/go-rest-template/pkg/models"
)

// GetAllBooks retrieves all books
func GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	result := db.DB.Find(&books)
	return books, result.Error
}

// GetBookByID retrieves a book by ID
func GetBookByID(id uint) (models.Book, error) {
	var book models.Book
	result := db.DB.First(&book, id)
	return book, result.Error
}

// CreateBook creates a new book
func CreateBook(book *models.Book) error {
	result := db.DB.Create(book)
	return result.Error
}

// UpdateBook updates an existing book
func UpdateBook(book *models.Book) error {
	result := db.DB.Save(book)
	return result.Error
}

// DeleteBook deletes a book by ID
func DeleteBook(id uint) error {
	result := db.DB.Delete(&models.Book{}, id)
	return result.Error
}
