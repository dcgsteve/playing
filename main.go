package main

import (
	"github.com/gin-gonic/gin"

	"github.com/dcgsteve/playing/controllers"
	"github.com/dcgsteve/playing/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/health", controllers.Health)

	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run()
}
