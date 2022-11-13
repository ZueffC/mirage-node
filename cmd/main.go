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

	packages := router.Group("/packages")
	{
		packages.POST("/", controllers.PackagesController)
	}

	router.Run()
}
