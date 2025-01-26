package routes

import (
	"library-crud-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/books", controllers.GetBooks)
	r.POST("/books", controllers.CreateBook)

	r.GET("/authors", controllers.GetAuthors)
	r.POST("/authors", controllers.CreateAuthor)

	return r
}
