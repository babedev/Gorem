package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/babedev/gorem/models"
)

func main() {
	p := os.Getenv("PORT")

	if p == "" {
		p = "8080"
	}

	r :=  gin.New()
	r.Use(gin.Logger())

	r.GET("/users", getUsers)

	r.Run(":" + p)
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, models.Users())
}