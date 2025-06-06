package main

import (
	"fmt"

	"blog-api/handlers"
	"blog-api/middleware"
	"blog-api/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize services
	postService := services.NewPostService("./posts")

	// Initialize handlers
	postHandler := handlers.NewPostHandler(postService)

	// Setup router
	r := gin.Default()

	// Add middleware
	r.Use(middleware.CORS())

	// Routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Blog API is running!",
			"endpoints": gin.H{
				"GET /posts":       "List all blog posts",
				"GET /posts/:slug": "Get a specific blog post",
				"GET /rss":         "RSS feed",
			},
		})
	})

	r.GET("/posts", postHandler.GetAllPosts)
	r.GET("/posts/:slug", postHandler.GetPostBySlug)
	r.GET("/rss", postHandler.GetRSSFeed)

	fmt.Println("Blog API starting on port 8080...")
	fmt.Println("Endpoints:")
	fmt.Println("  GET /        - API info")
	fmt.Println("  GET /posts   - List all posts")
	fmt.Println("  GET /posts/:slug - Get specific post")
	fmt.Println("  GET /rss     - RSS feed")

	r.Run(":8080")
}
