package main

import "gopkg.in/gin-gonic/gin.v1"
import "net/http"
import "github.com/babedev/gorem/models"

func main() {
	r :=  gin.New()
	r.Use(gin.Logger())

	r.GET("/users", getUsers)

	r.Run(":8080")
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, models.Users())
}