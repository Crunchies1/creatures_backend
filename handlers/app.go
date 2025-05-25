package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/Crunchies1/creatures_backend/internal"
	"github.com/Crunchies1/creatures_backend/models"
)

const (
	baseAPIPath = "api/v1"
)

type App struct {
	Router      *gin.Engine
	Client      *models.Client
	UserService *internal.UserService
}

func New(router *gin.Engine, client *models.Client) *App {
	app := &App{
		Router:      router,
		Client:      client,
		UserService: internal.NewUserService(client),
	}
	return app
}

func (app *App) SetupRoutes(router *gin.Engine) error {
	apiVersion := router.Group(baseAPIPath)
	app.registerUserRoutes(apiVersion)
	return nil
}
