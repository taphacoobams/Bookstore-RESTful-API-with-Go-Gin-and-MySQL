package handlers

import (
	"example/bookstore/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Homepage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Bookstore RESTful API with Go, Gin and MySQL"})
}

func GetBooks(c *gin.Context) {
	books, err := models.GetBooks()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, books)
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")

	book, err := models.GetBookByID(id)
	if err != nil {
		if err.Error() == "book not found" {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func PostBooks(c *gin.Context) {
	var newBook models.Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	if err := newBook.AddBook(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Book has been successfully added"})
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var updatedBook models.Book

	if err := c.BindJSON(&updatedBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid book data"})
		return
	}

	if err := models.UpdateBook(id, updatedBook); err != nil {
		if err.Error() == "book not found" {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error updating book: " + err.Error()})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Book has been successfully updated"})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	if err := models.DeleteBook(id); err != nil {
		if err.Error() == "book not found" {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error deleting book: " + err.Error()})
		}
		return
	}
	c.IndentedJSON(http.StatusNoContent, nil)
}
