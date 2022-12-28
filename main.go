package main

import (
	"AnarchyStock/database"
	"AnarchyStock/handlers"
	"AnarchyStock/models"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.InitDB()

	// database.DB.DB.AutoMigrate(models.User{})
	database.DB.DB.AutoMigrate(models.Product{})
	database.DB.DB.AutoMigrate(models.Order{})

	r := gin.Default()
	if os.Getenv("PRODUCTION") == "true" {
		config := cors.DefaultConfig()
		config.AllowOrigins = []string{
			os.Getenv("DOMAIN"),
		}
		r.Use(cors.New(config))
		r.SetTrustedProxies(nil)
	} else {
		config := cors.DefaultConfig()
		config.AllowOrigins = []string{
			"http://127.0.0.1:5173",
			"*",
		}
		r.Use(cors.New(config))
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// api := r.Group("/api")

	// authApi := api.Group("/auth")
	// authApi.POST("/login", handlers.Login)
	// authApi.POST("/sign-up", handlers.Register)

	// Protected routes
	// api.Use(middlewares.JwtAuthMiddleware())

	// Only protected routes

	// Protected Routes with user valid operation on the same user that is logged in
	// api.GET("/:username", handlers.CurrentUser)
	// products := api.Group("/products")
	r.GET("/products", handlers.GetAllProducts)
	r.GET("/products/paginated/:page", handlers.GetProductsPaginated)
	r.POST("/products/create", handlers.CreateProduct)
	r.POST("/products/search", handlers.SearchProduct)
	r.Run() // listen and serve on 0.0.0.0:8080
}
