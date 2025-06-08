package models

import "time"

// Custom time type that serializes as date only
type DateOnly time.Time

// MarshalJSON implements the json.Marshaler interface
func (d DateOnly) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(d).Format("2006-01-02") + `"`), nil
}

// IsZero reports whether d represents the zero time instant
func (d DateOnly) IsZero() bool {
	return time.Time(d).IsZero()
}

// Equal reports whether d and u represent the same time instant
func (d DateOnly) Equal(u DateOnly) bool {
	return time.Time(d).Equal(time.Time(u))
}

// BlogPost represents a blog post with metadata
type BlogPost struct {
	Slug        string    `json:"slug" example:"hello-world"`
	Title       string    `json:"title" example:"Hello World"`
	Date        DateOnly  `json:"date" example:"2024-01-01"`
	Tags        []string  `json:"tags,omitempty" example:"go,api,blog"`
	Content     string    `json:"content" example:"This is the full content of the blog post..."`
	Excerpt     string    `json:"excerpt,omitempty" example:"This is a short excerpt..."`
	PublishDate string    `json:"publish_date" example:"2024-01-01T12:00:00Z"`
}

// BlogPostMeta represents blog post metadata without content
type BlogPostMeta struct {
	Slug        string    `json:"slug" example:"hello-world"`
	Title       string    `json:"title" example:"Hello World"`
	Date        DateOnly  `json:"date" example:"2024-01-01"`
	Tags        []string  `json:"tags,omitempty" example:"go,api,blog"`
	Excerpt     string    `json:"excerpt,omitempty" example:"This is a short excerpt..."`
	PublishDate string    `json:"publish_date" example:"2024-01-01T12:00:00Z"`
}

// HealthResponse represents the health check response
type HealthResponse struct {
	Status    string `json:"status" example:"healthy"`
	Timestamp string `json:"timestamp" example:"2024-01-01T12:00:00Z"`
	Uptime    string `json:"uptime" example:"1h23m45s"`
	Service   string `json:"service" example:"blog-api"`
	Version   string `json:"version" example:"1.0.0"`
}

// ReadinessResponse represents the readiness check response
type ReadinessResponse struct {
	Status    string `json:"status" example:"ready"`
	Timestamp string `json:"timestamp" example:"2024-01-01T12:00:00Z"`
}

// LivenessResponse represents the liveness check response
type LivenessResponse struct {
	Status    string `json:"status" example:"alive"`
	Timestamp string `json:"timestamp" example:"2024-01-01T12:00:00Z"`
}

// PostsResponse represents the response for getting all posts
type PostsResponse struct {
	Posts []BlogPostMeta `json:"posts"`
	Count int            `json:"count" example:"5"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error" example:"Something went wrong"`
}
