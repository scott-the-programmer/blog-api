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
	Slug        string    `json:"slug"`
	Title       string    `json:"title"`
	Date        DateOnly  `json:"date"`
	Tags        []string  `json:"tags,omitempty"`
	Content     string    `json:"content"`
	Excerpt     string    `json:"excerpt,omitempty"`
	PublishDate string    `json:"publish_date"`
}

// BlogPostMeta represents blog post metadata without content
type BlogPostMeta struct {
	Slug        string    `json:"slug"`
	Title       string    `json:"title"`
	Date        DateOnly  `json:"date"`
	Tags        []string  `json:"tags,omitempty"`
	Excerpt     string    `json:"excerpt,omitempty"`
	PublishDate string    `json:"publish_date"`
}
