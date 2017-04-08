package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/babedev/gorem/models"
)

func List(c *gin.Context) {
	us := []models.User{
		{Fullname: "John Doe", Firstname: "John", Lastname: "Doe", Age: 15},
		{Fullname: "John Doe", Firstname: "John", Lastname: "Doe", Age: 15},
	}

	c.JSON(http.StatusOK, us)
}

func Get(c *gin.Context) {

	//c.JSON(http.StatusOK, user)
}
