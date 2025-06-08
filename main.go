package main

import (
	"fmt"

	"blog-api/handlers"
	"blog-api/middleware"
	"blog-api/services"

	"github.com/gin-gonic/gin"
)

func main() {
	postService := services.NewPostService("./posts")

	postHandler := handlers.NewPostHandler(postService)
	healthHandler := handlers.NewHealthHandler()

	r := gin.Default()

	r.Use(middleware.CORS())

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

	r.Run(":8080")
}
