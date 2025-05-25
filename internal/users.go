package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Crunchies1/creatures_backend/models"
)

var fake_users = []models.User{
	{ID: "1", Username: "John Doe", Email: "john.doe@example.com", Password: "password123"},
	{ID: "2", Username: "Jane Smith", Email: "jane.smith@example.com", Password: "password456"},
}

func CreateUser(c *gin.Context) {
	var newUser models.User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	fake_users = append(fake_users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, fake_users)
}
