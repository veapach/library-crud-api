package controllers

import (
	"library-crud-api/internal/database"
	"library-crud-api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	database.DB.Preload("Author").Find(&books)
	c.JSON(http.StatusOK, books)
}

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&book)
	c.JSON(http.StatusOK, book)
}

func GetBookById(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if err := database.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)

}

func DeleteBookById(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if err := database.DB.Delete(&book, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
