package models

import "time"

// BlogPost represents a blog post with metadata
type BlogPost struct {
	Slug        string    `json:"slug"`
	Title       string    `json:"title"`
	Date        time.Time `json:"date"`
	Tags        []string  `json:"tags,omitempty"`
	Content     string    `json:"content"`
	Excerpt     string    `json:"excerpt,omitempty"`
	PublishDate string    `json:"publish_date"`
}

// BlogPostMeta represents blog post metadata without content
type BlogPostMeta struct {
	Slug        string    `json:"slug"`
	Title       string    `json:"title"`
	Date        time.Time `json:"date"`
	Tags        []string  `json:"tags,omitempty"`
	Excerpt     string    `json:"excerpt,omitempty"`
	PublishDate string    `json:"publish_date"`
}
