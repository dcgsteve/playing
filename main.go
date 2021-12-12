package main

import (
	"github.com/gin-gonic/gin"

	"github.com/dcgsteve/playing/controllers" // new
	"github.com/dcgsteve/playing/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks) // new

	r.Run()
}
