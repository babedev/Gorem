package main

import (
	"os"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/babedev/gorem/handlers/users"
)

func main() {
	p := os.Getenv("PORT")

	if p == "" {
		p = "8080"
	}

	r := gin.New()
	r.LoadHTMLGlob("templates/*")
	r.Use(gin.Logger())

	r.GET("/", index)
	r.GET("/users", users.List)
	r.GET("/users/:id", users.Get)

	r.Run(":" + p)
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Directories",
	})
}
