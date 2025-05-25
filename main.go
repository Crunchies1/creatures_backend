package main

import (
	"github.com/gin-gonic/gin"

	"github.com/Crunchies1/creatures_backend/internal"
)

func main() {
	router := gin.Default()
	router.GET("/users", internal.GetUsers)
	router.POST("/users", internal.CreateUser)

	router.Run("localhost:8080")
}
