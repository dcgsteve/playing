package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /books
// Get all books
func Health(c *gin.Context) {
	c.HTML(http.StatusOK, "<html>OK</html>", gin.H{"title": "health check"})
}
