package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Crunchies1/creatures_backend/models"
)

func (a *App) registerUserRoutes(apiVersion *gin.RouterGroup) {
	users := apiVersion.Group("/users")
	users.POST("", a.createUser)     // POST /users
	users.GET("", a.getUsers)        // GET /users
	users.GET("/:id", a.getUserByID) // GET /users/:id
}

func (a *App) createUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := a.UserService.CreateUser(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newUser)
}

func (a *App) getUsers(c *gin.Context) {
	users, err := a.UserService.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (a *App) getUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := a.UserService.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
