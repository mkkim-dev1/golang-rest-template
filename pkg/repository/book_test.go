package repository

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"serengeti.app/go-rest-template/pkg/db"
	"serengeti.app/go-rest-template/pkg/models"
)

// TestMain sets up the in-memory database for testing
func setupTestDB() {
	var err error
	db.DB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.DB.AutoMigrate(&models.Book{})
}

func TestCreateBook(t *testing.T) {
	setupTestDB()

	book := models.Book{
		Title:  "Test Book",
		Author: "Test Author",
	}

	if err := CreateBook(&book); err != nil {
		t.Fatalf("Failed to create book: %v", err)
	}

	if book.ID == 0 {
		t.Errorf("Expected book ID to be set, but got 0")
	}
}

func TestGetBookByID(t *testing.T) {
	setupTestDB()

	// Create a book first
	book := models.Book{Title: "Test Book", Author: "Test Author"}
	CreateBook(&book)

	retrievedBook, err := GetBookByID(book.ID)
	if err != nil {
		t.Fatalf("Failed to get book by ID: %v", err)
	}

	if retrievedBook.Title != book.Title || retrievedBook.Author != book.Author {
		t.Errorf("Expected book title %s and author %s, but got %s and %s",
			book.Title, book.Author, retrievedBook.Title, retrievedBook.Author)
	}
}

func TestGetAllBooks(t *testing.T) {
	setupTestDB()

	books := []models.Book{
		{Title: "Book 1", Author: "Author 1"},
		{Title: "Book 2", Author: "Author 2"},
	}

	for _, book := range books {
		CreateBook(&book)
	}

	allBooks, err := GetAllBooks()
	if err != nil {
		t.Fatalf("Failed to retrieve all books: %v", err)
	}

	if len(allBooks) != len(books) {
		t.Errorf("Expected %d books, but got %d", len(books), len(allBooks))
	}
}

func TestUpdateBook(t *testing.T) {
	setupTestDB()

	// Create a book first
	book := models.Book{Title: "Old Title", Author: "Old Author"}
	CreateBook(&book)

	// Update the book
	book.Title = "New Title"
	book.Author = "New Author"
	if err := UpdateBook(&book); err != nil {
		t.Fatalf("Failed to update book: %v", err)
	}

	updatedBook, _ := GetBookByID(book.ID)
	if updatedBook.Title != "New Title" || updatedBook.Author != "New Author" {
		t.Errorf("Expected updated title 'New Title' and author 'New Author', but got %s and %s",
			updatedBook.Title, updatedBook.Author)
	}
}

func TestDeleteBook(t *testing.T) {
	setupTestDB()

	// Create a book first
	book := models.Book{Title: "Delete Me", Author: "Delete Author"}
	CreateBook(&book)

	if err := DeleteBook(book.ID); err != nil {
		t.Fatalf("Failed to delete book: %v", err)
	}

	_, err := GetBookByID(book.ID)
	if err == nil {
		t.Errorf("Expected error when retrieving deleted book, but got nil")
	}
}
