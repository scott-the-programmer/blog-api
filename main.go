package main

import (
	"fmt"

	"blog-api/handlers"
	"blog-api/middleware"
	"blog-api/services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "blog-api/docs"
)

// @title Blog API
// @version 1.0
// @description A simple blog API built with Go and Gin
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host https://blog-api.murray.kiwi
// @BasePath /

func main() {
	postService := services.NewPostService("./posts")

	postHandler := handlers.NewPostHandler(postService)
	healthHandler := handlers.NewHealthHandler()

	r := gin.Default()

	r.Use(middleware.CORS())

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// @Summary API Information
	// @Description Get basic information about the Blog API and available endpoints
	// @Tags general
	// @Accept json
	// @Produce json
	// @Success 200 {object} map[string]interface{} "message and endpoints list"
	// @Router / [get]
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Blog API is running!",
			"endpoints": gin.H{
				"GET /posts":       "List all blog posts",
				"GET /posts/:slug": "Get a specific blog post",
				"GET /rss":         "RSS feed",
				"GET /health":      "Health check",
				"GET /health/ready": "Readiness check",
				"GET /health/live":  "Liveness check",
				"GET /swagger/":     "API documentation",
			},
		})
	})

	// Health check endpoints
	r.GET("/health", healthHandler.HealthCheck)
	r.GET("/health/ready", healthHandler.ReadinessCheck)
	r.GET("/health/live", healthHandler.LivenessCheck)

	r.GET("/posts", postHandler.GetAllPosts)
	r.GET("/posts/", func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "Post not found"})
	})
	r.GET("/posts/:slug", postHandler.GetPostBySlug)
	r.GET("/rss", postHandler.GetRSSFeed)

	fmt.Println("Blog API starting on port 8080...")
	fmt.Println("Endpoints:")
	fmt.Println("  GET /        - API info")
	fmt.Println("  GET /health  - Health check")
	fmt.Println("  GET /health/ready - Readiness check")
	fmt.Println("  GET /health/live  - Liveness check")
	fmt.Println("  GET /posts   - List all posts")
	fmt.Println("  GET /posts/:slug - Get specific post")
	fmt.Println("  GET /rss     - RSS feed")
	fmt.Println("  GET /swagger/ - API documentation")

	r.Run(":8080")
}
