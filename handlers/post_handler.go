package handlers

import (
	"blog-api/models"
	"blog-api/services"

	"github.com/gin-gonic/gin"
)

// PostHandler handles HTTP requests for blog posts
type PostHandler struct {
	postService *services.PostService
}

// NewPostHandler creates a new PostHandler instance
func NewPostHandler(postService *services.PostService) *PostHandler {
	return &PostHandler{
		postService: postService,
	}
}

// GetAllPosts returns a list of all blog posts (without full content)
func (ph *PostHandler) GetAllPosts(c *gin.Context) {
	posts, err := ph.postService.GetAllPosts(false)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to load posts: " + err.Error()})
		return
	}

	// Convert to metadata only
	var postMetas []models.BlogPostMeta
	for _, post := range posts {
		postMetas = append(postMetas, models.BlogPostMeta{
			Slug:        post.Slug,
			Title:       post.Title,
			Date:        post.Date,
			Tags:        post.Tags,
			Excerpt:     post.Excerpt,
			PublishDate: post.PublishDate,
		})
	}

	c.JSON(200, gin.H{
		"posts": postMetas,
		"count": len(postMetas),
	})
}

// GetPostBySlug returns a specific blog post by its slug
func (ph *PostHandler) GetPostBySlug(c *gin.Context) {
	slug := c.Param("slug")

	post, err := ph.postService.GetPostBySlug(slug)
	if err != nil {
		c.JSON(404, gin.H{"error": "Post not found: " + err.Error()})
		return
	}

	c.JSON(200, post)
}

// GetRSSFeed returns an RSS feed of blog posts
func (ph *PostHandler) GetRSSFeed(c *gin.Context) {
	// You can make these configurable via environment variables or config file
	title := "My Blog"
	baseURL := "http://localhost:8080" // Change this to your actual domain
	description := "Latest posts from my blog"

	feed, err := ph.postService.GenerateRSSFeed(title, baseURL, description)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate RSS feed: " + err.Error()})
		return
	}

	c.Header("Content-Type", "application/rss+xml; charset=utf-8")
	c.String(200, feed.ToXML())
}
