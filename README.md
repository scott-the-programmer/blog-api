# Blog API

A simple Go API built with Gin that serves blog posts from markdown files. Stupidly simple, just the way it should be!

## Features

- 📝 Write blog posts in Markdown
- 🏷️ YAML frontmatter for metadata (title, date, tags, excerpt)
- 🚀 Fast and lightweight Go API using Gin
- 📱 JSON API responses perfect for frontends
- 🔄 CORS enabled for web applications
- 📂 File-based storage (no database needed)

## Quick Start

1. **Clone and setup:**

   ```bash
   git clone <your-repo>
   cd blog-api
   go mod tidy
   ```

2. **Run the API:**

   ```bash
   go run main.go
   ```

3. **API will be available at:** `http://localhost:8080`

## API Endpoints

| Method | Endpoint       | Description                              |
| ------ | -------------- | ---------------------------------------- |
| GET    | `/`            | API information and available endpoints  |
| GET    | `/posts`       | List all blog posts (metadata only)      |
| GET    | `/posts/:slug` | Get specific blog post with full content |

## Writing Blog Posts

Create markdown files in the `posts/` directory. Each file should have:

1. **Filename**: Use kebab-case, this becomes the URL slug (e.g., `my-first-post.md` → `/posts/my-first-post`)

2. **Frontmatter**: YAML metadata at the top:

   ```yaml
   ---
   title: "Your Post Title"
   date: "2025-06-05"
   tags: ["tag1", "tag2", "tag3"]
   excerpt: "A brief description of your post"
   ---
   ```

3. **Content**: Regular markdown below the frontmatter

### Example Post

````markdown
---
title: "My Awesome Post"
date: "2025-06-05"
tags: ["golang", "api", "blogging"]
excerpt: "This is what my post is about"
---

# My Awesome Post

This is the content of my post written in **markdown**.

## Subheading

- List item 1
- List item 2

`code blocks work too`
````

## Example API Responses

### GET /posts

```json
{
  "posts": [
    {
      "slug": "welcome-to-my-blog",
      "title": "Welcome to My Blog",
      "date": "2025-06-05T00:00:00Z",
      "tags": ["welcome", "first-post", "blogging"],
      "excerpt": "This is my first blog post using the new markdown-based blog API!",
      "publish_date": "2025-06-05"
    }
  ],
  "count": 1
}
```

### GET /posts/welcome-to-my-blog

```json
{
  "slug": "welcome-to-my-blog",
  "title": "Welcome to My Blog",
  "date": "2025-06-05T00:00:00Z",
  "tags": ["welcome", "first-post", "blogging"],
  "content": "# Welcome to My Blog\n\nThis is my first blog post...",
  "excerpt": "This is my first blog post using the new markdown-based blog API!",
  "publish_date": "2025-06-05"
}
```

## Development

- Posts are automatically sorted by date (newest first)
- If no excerpt is provided, one is auto-generated from the content
- If no date is provided, the file modification time is used
- CORS is enabled for all origins (configure as needed for production)

## Deployment

This is a standard Go application. You can:

1. **Build binary:**

   ```bash
   go build -o blog-api main.go
   ./blog-api
   ```

2. **Docker:** (create a Dockerfile as needed)

3. **Deploy to any platform** that supports Go applications

## License

MIT License - feel free to use this for your own blog!
