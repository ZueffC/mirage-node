package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zueffc/mirage-node/internal/controllers"
	"github.com/zueffc/mirage-node/internal/data"
)

func main() {
	router := gin.Default()

	data.Connect()

	users := router.Group("/users")
	{
		users.POST("/auth", controllers.AuthController)
		users.POST("/delete", controllers.DeleteController)
	}

	router.POST("pkgs/search", nil)
	router.POST("pkgs/add", nil)
	router.POST("pkgs/edit")

	router.Run()
}
