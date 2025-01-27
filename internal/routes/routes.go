package routes

import (
	"library-crud-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//CRUD FOR BOOKS
	r.GET("/books", controllers.GetBooks)
	r.GET("/books/:id", controllers.GetBookById)
	r.POST("/books", controllers.CreateBook)
	r.DELETE("/books/:id", controllers.DeleteBookById)

	//CRUD FOR AUTHORS
	r.GET("/authors", controllers.GetAuthors)
	r.POST("/authors", controllers.CreateAuthor)

	return r
}
