package main

import (
	"library-crud-api/internal/database"
	"library-crud-api/internal/routes"
)

func main() {
	database.Connect()

	r := routes.SetupRouter()
	r.Run(":8080")
}
