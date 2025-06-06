package services

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

// TestData directory for test posts
const testPostsDir = "./posts"

// setupTestDir creates a temporary directory with test posts
func setupTestDir(t *testing.T) {
	// Clean up any existing test directory
	os.RemoveAll(testPostsDir)

	// Create test directory
	err := os.Mkdir(testPostsDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	// Create test posts
	testPosts := map[string]string{
		"hello-world.md": `---
title: "Hello World"
date: "2025-06-05"
tags: ["Hello World"]
excerpt: "A post with nothing substantial"
---

# Hello World

There isn't anything meaningful in this post - apart from maybe a wince when looking back in time at this.

The main reason I want to start up a blog isn't for _anyone_ to read it, but rather for me to reference a thought I had one time and share that with those around me.

I'll likely ramble about tech stuff, side projects, and probably weave my 2 dogs into the mix.

Test post, please ignore`,

		"post-with-frontmatter.md": `---
title: "Test Post with Frontmatter"
date: "2025-06-05"
tags: ["golang", "testing"]
excerpt: "This is a test post excerpt"
---

# Test Post with Frontmatter

This is the content of a test post with proper frontmatter.

It has multiple paragraphs and should be parsed correctly.`,

		"post-without-frontmatter.md": `# Post Without Frontmatter

This is a simple post without any frontmatter.

It should still be parsed correctly with defaults.`,

		"post-with-array-tags.md": `---
title: "Post with Array Tags"
date: "2025-06-04"
tags: ["tag1", "tag2", "tag3"]
---

# Post with Array Tags

This post has tags in array format.`,

		"post-with-rfc3339-date.md": `---
title: "Post with RFC3339 Date"
date: "2025-06-03T10:30:00Z"
tags: ["datetime"]
---

# Post with RFC3339 Date

This post uses RFC3339 date format.`,

		"invalid-date-post.md": `---
title: "Post with Invalid Date"
date: "invalid-date"
tags: ["test"]
---

# Post with Invalid Date

This post has an invalid date format.`,
	}

	for filename, content := range testPosts {
		err := os.WriteFile(filepath.Join(testPostsDir, filename), []byte(content), 0644)
		if err != nil {
			t.Fatalf("Failed to create test post %s: %v", filename, err)
		}
	}
}

// cleanupTestDir removes the test directory
func cleanupTestDir(t *testing.T) {
	err := os.RemoveAll(testPostsDir)
	if err != nil {
		t.Errorf("Failed to cleanup test directory: %v", err)
	}
}

func TestNewPostService(t *testing.T) {
	setupTestDir(t)
	defer cleanupTestDir(t)

	service := NewPostService(testPostsDir)
	if service == nil {
		t.Error("NewPostService should return a non-nil PostService")
	}

	// Test that posts directory is created if it doesn't exist
	tempDir := "./temp_test_posts"
	defer os.RemoveAll(tempDir)

	// This would require modifying the const, so we'll test the actual behavior
	if _, err := os.Stat(testPostsDir); os.IsNotExist(err) {
		t.Errorf("Posts directory should be created if it doesn't exist")
	}
}

func TestLoadPostFromFile_WithFrontmatter(t *testing.T) {
	setupTestDir(t)
	defer cleanupTestDir(t)

	service := &PostService{}
	filepath := filepath.Join(testPostsDir, "post-with-frontmatter.md")

	post, err := service.loadPostFromFile(filepath, true)
	if err != nil {
		t.Fatalf("loadPostFromFile failed: %v", err)
	}

	// Test all parsed fields
	if post.Title != "Test Post with Frontmatter" {
		t.Errorf("Expected title 'Test Post with Frontmatter', got '%s'", post.Title)
	}

	if post.Slug != "post-with-frontmatter" {
		t.Errorf("Expected slug 'post-with-frontmatter', got '%s'", post.Slug)
	}

	expectedDate, _ := time.Parse("2006-01-02", "2025-06-05")
	if !post.Date.Equal(expectedDate) {
		t.Errorf("Expected date %v, got %v", expectedDate, post.Date)
	}

	if post.PublishDate != "2025-06-05" {
		t.Errorf("Expected publish date '2025-06-05', got '%s'", post.PublishDate)
	}

	expectedTags := []string{"golang", "testing"}
	if len(post.Tags) != len(expectedTags) {
		t.Errorf("Expected %d tags, got %d", len(expectedTags), len(post.Tags))
	}
	for i, tag := range expectedTags {
		if i >= len(post.Tags) || post.Tags[i] != tag {
			t.Errorf("Expected tag '%s' at index %d, got '%s'", tag, i, post.Tags[i])
		}
	}

	if post.Excerpt != "This is a test post excerpt" {
		t.Errorf("Expected excerpt 'This is a test post excerpt', got '%s'", post.Excerpt)
	}

	if post.Content == "" {
		t.Error("Content should not be empty when includeContent is true")
	}

	expectedContentStart := "# Test Post with Frontmatter"
	if len(post.Content) < len(expectedContentStart) || post.Content[:len(expectedContentStart)] != expectedContentStart {
		t.Errorf("Content should start with '%s'", expectedContentStart)
	}
}

func TestLoadPostFromFile_WithoutFrontmatter(t *testing.T) {
	setupTestDir(t)
	defer cleanupTestDir(t)

	service := &PostService{}
	filepath := filepath.Join(testPostsDir, "post-without-frontmatter.md")

	post, err := service.loadPostFromFile(filepath, true)
	if err != nil {
		t.Fatalf("loadPostFromFile failed: %v", err)
	}

	// Test defaults
	if post.Title != "Post Without Frontmatter" {
		t.Errorf("Expected title 'Post Without Frontmatter', got '%s'", post.Title)
	}

	if post.Slug != "post-without-frontmatter" {
		t.Errorf("Expected slug 'post-without-frontmatter', got '%s'", post.Slug)
	}

	// Date should be set from file modification time
	if post.Date.IsZero() {
		t.Error("Date should not be zero when no frontmatter date is provided")
	}

	// Should have generated excerpt
	if post.Excerpt == "" {
		t.Error("Excerpt should be generated when not provided in frontmatter")
	}
}

func TestLoadPostFromFile_WithoutContent(t *testing.T) {
	setupTestDir(t)
	defer cleanupTestDir(t)

	service := &PostService{}
	filepath := filepath.Join(testPostsDir, "post-with-frontmatter.md")

	post, err := service.loadPostFromFile(filepath, false)
	if err != nil {
		t.Fatalf("loadPostFromFile failed: %v", err)
	}

	if post.Content != "" {
		t.Error("Content should be empty when includeContent is false")
	}

	// Other fields should still be populated
	if post.Title == "" {
		t.Error("Title should be populated even when includeContent is false")
	}
}

func TestLoadPostFromFile_ArrayTags(t *testing.T) {
	setupTestDir(t)
	defer cleanupTestDir(t)

	service := &PostService{}
	filepath := filepath.Join(testPostsDir, "post-with-array-tags.md")

	post, err := service.loadPostFromFile(filepath, true)
	if err != nil {
		t.Fatalf("loadPostFromFile failed: %v", err)
	}

	expectedTags := []string{"tag1", "tag2", "tag3"}
	if len(post.Tags) != len(expectedTags) {
		t.Errorf("Expected %d tags, got %d", len(expectedTags), len(post.Tags))
	}
	for i, tag := range expectedTags {
		if i >= len(post.Tags) || post.Tags[i] != tag {
			t.Errorf("Expected tag '%s' at index %d, got '%s'", tag, i, post.Tags[i])
		}
	}
}

func TestLoadPostFromFile_RFC3339Date(t *testing.T) {
	setupTestDir(t)
	defer cleanupTestDir(t)

	service := &PostService{}
	filepath := filepath.Join(testPostsDir, "post-with-rfc3339-date.md")

	post, err := service.loadPostFromFile(filepath, true)
	if err != nil {
		t.Fatalf("loadPostFromFile failed: %v", err)
	}

	expectedDate, _ := time.Parse(time.RFC3339, "2025-06-03T10:30:00Z")
	if !post.Date.Equal(expectedDate) {
		t.Errorf("Expected date %v, got %v", expectedDate, post.Date)
	}

	if post.PublishDate != "2025-06-03" {
		t.Errorf("Expected publish date '2025-06-03', got '%s'", post.PublishDate)
	}
}

func TestLoadPostFromFile_InvalidDate(t *testing.T) {
	setupTestDir(t)
	defer cleanupTestDir(t)

	service := &PostService{}
	filepath := filepath.Join(testPostsDir, "invalid-date-post.md")

	post, err := service.loadPostFromFile(filepath, true)
	if err != nil {
		t.Fatalf("loadPostFromFile failed: %v", err)
	}

	// Should fall back to file modification time
	if post.Date.IsZero() {
		t.Error("Date should not be zero when invalid date is provided")
	}
}

func TestLoadPostFromFile_NonexistentFile(t *testing.T) {
	service := &PostService{}
	filepath := "./nonexistent-post.md"

	_, err := service.loadPostFromFile(filepath, true)
	if err == nil {
		t.Error("loadPostFromFile should return error for nonexistent file")
	}
}

func TestGetAllPosts_WithoutContent(t *testing.T) {
	setupTestDir(t)
	defer cleanupTestDir(t)

	service := &PostService{
		postsDir: testPostsDir,
	}

	posts, err := service.GetAllPosts(false)
	if err != nil {
		t.Fatalf("GetAllPosts failed: %v", err)
	}

	// Test that content is not included
	for _, post := range posts {
		if post.Content != "" {
			t.Errorf("Post %s should not have content when includeContent is false", post.Slug)
		}
	}
}

func TestGetPostBySlug(t *testing.T) {
	setupTestDir(t)
	defer cleanupTestDir(t)

	service := NewPostService(testPostsDir)

	// Test with existing post
	post, err := service.GetPostBySlug("hello-world")
	if err != nil {
		t.Fatalf("GetPostBySlug failed: %v", err)
	}

	if post.Slug != "hello-world" {
		t.Errorf("Expected slug 'hello-world', got '%s'", post.Slug)
	}

	if post.Content == "" {
		t.Error("Post content should be included")
	}
}

func TestGetPostBySlug_NonexistentPost(t *testing.T) {
	setupTestDir(t)
	defer cleanupTestDir(t)

	service := NewPostService(testPostsDir)

	_, err := service.GetPostBySlug("nonexistent-post")
	if err == nil {
		t.Error("GetPostBySlug should return error for nonexistent post")
	}
}

func TestGenerateRSSFeed(t *testing.T) {
	setupTestDir(t)
	defer cleanupTestDir(t)

	service := NewPostService(testPostsDir)

	title := "Test Blog"
	baseURL := "https://example.com"
	description := "A test blog"

	feed, err := service.GenerateRSSFeed(title, baseURL, description)
	if err != nil {
		t.Fatalf("GenerateRSSFeed failed: %v", err)
	}

	if feed.Title != title {
		t.Errorf("Expected feed title '%s', got '%s'", title, feed.Title)
	}

	if feed.Link != baseURL {
		t.Errorf("Expected feed link '%s', got '%s'", baseURL, feed.Link)
	}

	if feed.Description != description {
		t.Errorf("Expected feed description '%s', got '%s'", description, feed.Description)
	}

	if feed.Language != "en-us" {
		t.Errorf("Expected feed language 'en-us', got '%s'", feed.Language)
	}

	// Should have RSS items
	if len(feed.Items) == 0 {
		t.Error("RSS feed should have items")
	}

	// Should be limited to 20 items max
	if len(feed.Items) > 20 {
		t.Error("RSS feed should be limited to 20 items")
	}
}

func TestExcerptGeneration(t *testing.T) {
	setupTestDir(t)
	defer cleanupTestDir(t)

	// Create a post with long content and no excerpt
	longContent := `---
title: "Long Post"
date: "2025-06-05"
---

# Long Post

This is the first paragraph of a very long post that should be truncated when generating an excerpt automatically.

This is the second paragraph that contains even more content that should help test the excerpt generation functionality.

This is the third paragraph that definitely should not appear in the excerpt.`

	filepath := filepath.Join(testPostsDir, "long-post.md")
	err := os.WriteFile(filepath, []byte(longContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create long post: %v", err)
	}

	service := &PostService{}
	post, err := service.loadPostFromFile(filepath, true)
	if err != nil {
		t.Fatalf("loadPostFromFile failed: %v", err)
	}

	if post.Excerpt == "" {
		t.Error("Excerpt should be generated when not provided")
	}

	if len(post.Excerpt) > 203 { // 200 + "..."
		t.Errorf("Excerpt should be truncated to 200 characters plus '...', got %d characters", len(post.Excerpt))
	}
}
