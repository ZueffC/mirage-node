package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"github.com/qor/admin"
	"github.com/zueffc/mirage-node/internal/controllers"
	"github.com/zueffc/mirage-node/internal/data"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	router := gin.Default()

	//connection to database
	data.Connect()

	mux := http.NewServeMux()

	Admin := admin.New(&admin.AdminConfig{DB: data.Database()})
	Admin.AddResource(&data.UserModel{})
	Admin.AddResource(&data.PackageModel{})
	Admin.MountTo("/admin", mux)

	users := router.Group("/users")
	{
		users.POST("/auth", controllers.AuthController)
		users.POST("/delete", controllers.DeleteController)
	}

	packages := router.Group("/packages")
	{
		packages.POST("/", controllers.PackagesController)
		packages.POST("/get", controllers.PackageInformationController)
	}

	router.Any("/admin/*resources", gin.WrapH(mux))

	sslManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("zueffc.ml"),
		Cache:      autocert.DirCache("/var/www/.cache"),
	}

	log.Fatal(autotls.RunWithManager(router, &sslManager))
}
