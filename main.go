package main

import (
	"fmt"
	"log"
	"my-module/controllers"
	"my-module/initializers"
	"my-module/models"
	"my-module/routes"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine

	ProductController      controllers.ProductController
	ProductRouteController routes.ProductRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	initializers.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	initializers.DB.AutoMigrate(&models.Product{}, &models.Brand{}, &models.Category{})
	fmt.Println("👍 Migration complete")

	ProductController = controllers.NewProductController(initializers.DB)
	ProductRouteController = routes.NewRouteProductController(ProductController)

	server = gin.Default()

}
func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthChecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	ProductRouteController.ProductRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))

}
