package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/babedev/gorem/models"
)

func List(c *gin.Context) {
	cs := models.Contact{Email: "john@doe.com", FacebookLink: "johndoe", TwitterLink: ""}
	us := []models.User{
		{ID: 1, Fullname: "John Doe", Firstname: "John", Lastname: "Doe", Age: 15, Contact: cs},
		{ID: 2, Fullname: "John Doe", Firstname: "John", Lastname: "Doe", Age: 15},
		{},
	}

	c.JSON(http.StatusOK, us)
}

func Get(c *gin.Context) {

	//c.JSON(http.StatusOK, user)
}
