package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"serengeti.app/go-rest-template/pkg/models"
	"serengeti.app/go-rest-template/pkg/repository"
)

func getBooks(c *gin.Context) {
	books, err := repository.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ApiResponse{Success: false, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.ApiResponse{Success: true, Content: books})
}

func getBookByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := repository.GetBookByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, models.ApiResponse{Success: false, Error: "Book not found"})
		return
	}
	c.JSON(http.StatusOK, models.ApiResponse{Success: true, Content: []models.Book{book}})
}

func createBook(c *gin.Context) {
	var book models.Book
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, models.ApiResponse{Success: false, Error: err.Error()})
		return
	}
	if err := repository.CreateBook(&book); err != nil {
		c.JSON(http.StatusInternalServerError, models.ApiResponse{Success: false, Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, models.ApiResponse{Success: true, Content: []models.Book{book}})
}

func updateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var book models.Book
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, models.ApiResponse{Success: false, Error: err.Error()})
		return
	}
	book.ID = uint(id)
	if err := repository.UpdateBook(&book); err != nil {
		c.JSON(http.StatusInternalServerError, models.ApiResponse{Success: false, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.ApiResponse{Success: true, Content: []models.Book{book}})
}

func deleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := repository.DeleteBook(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, models.ApiResponse{Success: false, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.ApiResponse{Success: true, Content: []models.Book{{
		ID:     uint(id),
		Title:  "",
		Author: "",
	}}})
}
