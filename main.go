package main

import (
	"log"
	"net/http"
	"os"

	"gopkg.in/gin-gonic/gin.v1"
	"github.com/babedev/gorem/models"
)

func main() {
	p := os.Getenv("PORT")

	if p == "" {
		log.Fatal("$PORT must be set")
	}

	r :=  gin.New()
	r.Use(gin.Logger())

	r.GET("/users", getUsers)

	r.Run(":" + p)
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, models.Users())
}