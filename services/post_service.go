package services

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"blog-api/models"
)

// PostService handles blog post operations
type PostService struct {
	postsDir string
}

// NewPostService creates a new PostService instance
func NewPostService(postsDir string) *PostService {
	if _, err := os.Stat(postsDir); os.IsNotExist(err) {
		os.Mkdir(postsDir, 0755)
	}
	return &PostService{
		postsDir: postsDir,
	}
}

// GetAllPosts loads all blog posts from the posts directory
func (ps *PostService) GetAllPosts(includeContent bool) ([]models.BlogPost, error) {
	var posts []models.BlogPost

	files, err := ioutil.ReadDir(ps.postsDir)
	if err != nil {
		return posts, err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".md" {
			post, err := ps.loadPostFromFile(filepath.Join(ps.postsDir, file.Name()), includeContent)
			if err != nil {
				fmt.Printf("Error loading post %s: %v\n", file.Name(), err)
				continue
			}
			posts = append(posts, post)
		}
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date)
	})

	return posts, nil
}

// GetPostBySlug loads a specific post by its slug
func (ps *PostService) GetPostBySlug(slug string) (models.BlogPost, error) {
	filename := slug + ".md"
	filePath := filepath.Join(ps.postsDir, filename)

	return ps.loadPostFromFile(filePath, true)
}

// loadPostFromFile loads a blog post from a markdown file
func (ps *PostService) loadPostFromFile(filePath string, includeContent bool) (models.BlogPost, error) {
	var post models.BlogPost

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return post, err
	}

	contentStr := string(content)

	frontmatterRegex := regexp.MustCompile(`(?s)^---\s*\n(.*?)\n---\s*\n(.*)`)
	matches := frontmatterRegex.FindStringSubmatch(contentStr)

	var frontmatter string
	var markdown string

	if len(matches) == 3 {
		frontmatter = matches[1]
		markdown = matches[2]
	} else {
		markdown = contentStr
	}

	post.Slug = strings.TrimSuffix(filepath.Base(filePath), ".md")

	if frontmatter != "" {
		lines := strings.Split(frontmatter, "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			parts := strings.SplitN(line, ":", 2)
			if len(parts) != 2 {
				continue
			}

			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			if (strings.HasPrefix(value, "\"") && strings.HasSuffix(value, "\"")) ||
				(strings.HasPrefix(value, "'") && strings.HasSuffix(value, "'")) {
				value = value[1 : len(value)-1]
			}

			switch key {
			case "title":
				post.Title = value
			case "date":
				if date, err := time.Parse("2006-01-02", value); err == nil {
					post.Date = date
					post.PublishDate = value
				} else if date, err := time.Parse(time.RFC3339, value); err == nil {
					post.Date = date
					post.PublishDate = date.Format("2006-01-02")
				}
			case "tags":
				if strings.HasPrefix(value, "[") && strings.HasSuffix(value, "]") {
					value = strings.Trim(value, "[]")
					tags := strings.Split(value, ",")
					for _, tag := range tags {
						tag = strings.TrimSpace(tag)
						tag = strings.Trim(tag, "\"'")
						if tag != "" {
							post.Tags = append(post.Tags, tag)
						}
					}
				} else {
					tags := strings.Split(value, ",")
					for _, tag := range tags {
						tag = strings.TrimSpace(tag)
						if tag != "" {
							post.Tags = append(post.Tags, tag)
						}
					}
				}
			case "excerpt":
				post.Excerpt = value
			}
		}
	}

	if post.Title == "" {
		post.Title = strings.Title(strings.ReplaceAll(post.Slug, "-", " "))
	}

	if post.Date.IsZero() {
		if info, err := os.Stat(filePath); err == nil {
			post.Date = info.ModTime()
			post.PublishDate = info.ModTime().Format("2006-01-02")
		}
	}

	if includeContent {
		post.Content = strings.TrimSpace(markdown)
	}

	if post.Excerpt == "" && includeContent {
		lines := strings.Split(markdown, "\n")
		var excerptLines []string
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line != "" && !strings.HasPrefix(line, "#") {
				excerptLines = append(excerptLines, line)
				if len(excerptLines) >= 2 {
					break
				}
			}
		}
		post.Excerpt = strings.Join(excerptLines, " ")
		if len(post.Excerpt) > 200 {
			post.Excerpt = post.Excerpt[:200] + "..."
		}
	}

	return post, nil
}

// GenerateRSSFeed creates an RSS feed from blog posts
func (ps *PostService) GenerateRSSFeed(title, baseURL, description string) (models.RSSFeed, error) {
	posts, err := ps.GetAllPosts(true)
	if err != nil {
		return models.RSSFeed{}, err
	}

	feed := models.RSSFeed{
		Title:       title,
		Link:        baseURL,
		Description: description,
		Language:    "en-us",
		Items:       make([]models.RSSItem, 0, len(posts)),
	}

	limit := len(posts)
	if limit > 20 {
		limit = 20
	}

	for i := 0; i < limit; i++ {
		item := models.BlogPostToRSSItem(posts[i], baseURL)
		feed.Items = append(feed.Items, item)
	}

	return feed, nil
}
