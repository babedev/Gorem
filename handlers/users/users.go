package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/babedev/gorem/models"
)

func List(c *gin.Context) {
	contact := models.Contact{Email: "john@doe.com", FacebookLink: "johndoe", TwitterLink: ""}
	result := []models.User{
		{ID: 1, Fullname: "John Doe", Firstname: "John", Lastname: "Doe", Age: 15, Contact: contact},
		{ID: 2, Fullname: "John Doe", Firstname: "John", Lastname: "Doe", Age: 15},
		{},
	}

	c.JSON(http.StatusOK, result)
}

func Get(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	contact := models.Contact{Email: "john@doe.com", FacebookLink: "johndoe", TwitterLink: ""}
	result := models.User{ID: uint(id), Fullname: "John Doe", Firstname: "John", Lastname: "Doe", Age: 15, Contact: contact}

	c.JSON(http.StatusOK, result)
}
