# Blog API

A simple Go API that serves blog posts from markdown files.

## Quick Start

```bash
make deps    # Install dependencies
make run     # Start the API at http://localhost:8080
```

## Makefile Commands

| Command              | Description                    |
| -------------------- | ------------------------------ |
| `make help`          | Show all available commands    |
| `make run`           | Run the application            |
| `make build`         | Build the application binary   |
| `make test`          | Run unit tests                 |
| `make test-e2e`      | Run end-to-end tests           |
| `make test-all`      | Run all tests                  |
| `make test-coverage` | Run tests with coverage report |
| `make fmt`           | Format code                    |
| `make vet`           | Vet code                       |
| `make lint`          | Run linter                     |
| `make clean`         | Clean build artifacts          |
| `make watch`         | Watch for changes and rebuild  |

## Writing Posts

Create `.md` files in the `posts/` directory with YAML frontmatter:

```yaml
---
title: "Your Post Title"
date: "2025-06-05"
tags: ["tag1", "tag2"]
excerpt: "Brief description"
---
# Your content here
```

## API Endpoints

- `GET /` - API info
- `GET /posts` - List all posts
- `GET /posts/:slug` - Get specific post
- `GET /health` - Health check
- `GET /rss` - RSS feed
