package main

import (
	"log"
	"time"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/sqlite"
	"github.com/gin-contrib/cors"
	"golang.org/x/crypto/acme/autocert"

	"github.com/GoAdminGroup/themes/adminlte"

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"github.com/zueffc/mirage-node/internal/controllers"
	"github.com/zueffc/mirage-node/internal/data"
	"github.com/zueffc/mirage-node/internal/data/foradmin"
)

func main() {
	router := gin.Default()

	eng := engine.Default()

	cfg := config.Config{
		Env: config.EnvLocal,
		Databases: config.DatabaseList{
			"default": {
				Host:            "localhost",
				Port:            "10000",
				User:            "root",
				Pwd:             "root",
				Name:            "godmin",
				MaxIdleConns:    50,
				MaxOpenConns:    150,
				ConnMaxLifetime: time.Hour,
				Driver:          config.DriverSqlite,
				File:            "./test_database.db",
			},
		},
		UrlPrefix:   "admin",
		ColorScheme: adminlte.ColorschemeSkinBlack,
	}

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://mirage-pm.ml"}

	router.Use(cors.New(config))

	//connection to database
	data.Connect()
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

	eng.AddConfig(&cfg).
		AddGenerators(foradmin.Generators).
		AddGenerator("package_models", foradmin.GetPackagesTable).
		Use(router)

	sslManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("zueffc.ml"),
		Cache:      autocert.DirCache("/var/www/.cache"),
	}

	log.Fatal(autotls.RunWithManager(router, &sslManager))
}
