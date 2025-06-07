package e2e

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// makeRequest is a helper function to make HTTP requests
func makeRequest(t *testing.T, method, url string) (*http.Response, []byte) {
	client := &http.Client{Timeout: 10 * time.Second}
	
	req, err := http.NewRequest(method, url, nil)
	require.NoError(t, err, "Failed to create request")

	resp, err := client.Do(req)
	require.NoError(t, err, "Failed to make request to %s", url)

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "Failed to read response body")
	defer resp.Body.Close()

	return resp, body
}

// getBaseURL returns the base URL for the API, either from environment or default
func getBaseURL() string {
	if url := os.Getenv("API_BASE_URL"); url != "" {
		return url
	}
	return "http://localhost:8080"
}

// checkServiceAvailability verifies the service is running before tests
func checkServiceAvailability(t *testing.T, baseURL string) {
	client := &http.Client{Timeout: 5 * time.Second}
	
	resp, err := client.Get(baseURL + "/health")
	if err != nil {
		t.Fatalf("Service is not running at %s. Please start the service first. Error: %v", baseURL, err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Service health check failed. Expected status 200, got %d", resp.StatusCode)
	}
	
	t.Logf("Service is running at %s", baseURL)
}

// setupTest checks that the service is available and returns the base URL
func setupTest(t *testing.T) string {
	baseURL := getBaseURL()
	checkServiceAvailability(t, baseURL)
	return baseURL
}

func TestBlogAPIEndpoints(t *testing.T) {
	baseURL := setupTest(t)

	t.Run("Root endpoint", func(t *testing.T) {
		resp, body := makeRequest(t, "GET", baseURL+"/")
		
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		
		var response map[string]interface{}
		err := json.Unmarshal(body, &response)
		require.NoError(t, err, "Failed to parse JSON response")
		
		assert.Equal(t, "Blog API is running!", response["message"])
		assert.Contains(t, response, "endpoints")
		
		endpoints, ok := response["endpoints"].(map[string]interface{})
		require.True(t, ok, "endpoints should be an object")
		
		// Verify expected endpoints are documented
		expectedEndpoints := []string{
			"GET /posts",
			"GET /posts/:slug", 
			"GET /rss",
			"GET /health",
			"GET /health/ready",
			"GET /health/live",
		}
		
		for _, endpoint := range expectedEndpoints {
			assert.Contains(t, endpoints, endpoint, "Endpoint %s should be documented", endpoint)
		}
	})

	t.Run("Health check endpoints", func(t *testing.T) {
		// Test main health endpoint
		t.Run("GET /health", func(t *testing.T) {
			resp, body := makeRequest(t, "GET", baseURL+"/health")
			
			assert.Equal(t, http.StatusOK, resp.StatusCode)
			
			var health map[string]interface{}
			err := json.Unmarshal(body, &health)
			require.NoError(t, err, "Failed to parse health response")
			
			assert.Equal(t, "healthy", health["status"])
			assert.Equal(t, "blog-api", health["service"])  
			assert.Equal(t, "1.0.0", health["version"])
			assert.Contains(t, health, "timestamp")
			assert.Contains(t, health, "uptime")
		})

		// Test readiness endpoint
		t.Run("GET /health/ready", func(t *testing.T) {
			resp, body := makeRequest(t, "GET", baseURL+"/health/ready")
			
			assert.Equal(t, http.StatusOK, resp.StatusCode)
			
			var readiness map[string]interface{}
			err := json.Unmarshal(body, &readiness)
			require.NoError(t, err, "Failed to parse readiness response")
			
			assert.Equal(t, "ready", readiness["status"])
			assert.Contains(t, readiness, "timestamp")
		})

		// Test liveness endpoint
		t.Run("GET /health/live", func(t *testing.T) {
			resp, body := makeRequest(t, "GET", baseURL+"/health/live")
			
			assert.Equal(t, http.StatusOK, resp.StatusCode)
			
			var liveness map[string]interface{}
			err := json.Unmarshal(body, &liveness)
			require.NoError(t, err, "Failed to parse liveness response")
			
			assert.Equal(t, "alive", liveness["status"])
			assert.Contains(t, liveness, "timestamp")
		})
	})

	t.Run("Post endpoints", func(t *testing.T) {
		// Test get all posts
		t.Run("GET /posts", func(t *testing.T) {
			resp, body := makeRequest(t, "GET", baseURL+"/posts")
			
			assert.Equal(t, http.StatusOK, resp.StatusCode)
			
			var response map[string]interface{}
			err := json.Unmarshal(body, &response)
			require.NoError(t, err, "Failed to parse posts response")
			
			assert.Contains(t, response, "posts")
			assert.Contains(t, response, "count")
			
			posts, ok := response["posts"].([]interface{})
			require.True(t, ok, "posts should be an array")
			
			count, ok := response["count"].(float64)
			require.True(t, ok, "count should be a number")
			
			assert.Equal(t, float64(len(posts)), count)
			
			// Verify we have at least the hello-world post
			assert.GreaterOrEqual(t, len(posts), 1, "Should have at least one post")
			
			// Check the structure of the first post
			if len(posts) > 0 {
				post := posts[0].(map[string]interface{})
				assert.Contains(t, post, "slug")
				assert.Contains(t, post, "title")  
				assert.Contains(t, post, "date")
				assert.Contains(t, post, "tags")
				assert.Contains(t, post, "excerpt")
				
				// Should not contain full content (content field)
				assert.NotContains(t, post, "content")
			}
		})

		// Test get specific post by slug
		t.Run("GET /posts/hello-world", func(t *testing.T) {
			resp, body := makeRequest(t, "GET", baseURL+"/posts/hello-world")
			
			assert.Equal(t, http.StatusOK, resp.StatusCode)
			
			var post map[string]interface{}
			err := json.Unmarshal(body, &post)
			require.NoError(t, err, "Failed to parse post response")
			
			// Verify post structure
			assert.Equal(t, "hello-world", post["slug"])
			assert.Equal(t, "Hello World", post["title"])
			assert.Equal(t, "2025-06-05", post["date"])
			assert.Contains(t, post, "tags")
			assert.Contains(t, post, "excerpt")
			assert.Contains(t, post, "content")
			
			// Verify tags structure
			tags, ok := post["tags"].([]interface{})
			require.True(t, ok, "tags should be an array")
			assert.Contains(t, tags, "Hello World")
			
			// Verify content is present (full post includes content)
			content, ok := post["content"].(string)
			require.True(t, ok, "content should be a string")
			assert.Contains(t, content, "Hello World")
			assert.Contains(t, content, "Test post, please ignore")
		})

		// Test get non-existent post
		t.Run("GET /posts/non-existent", func(t *testing.T) {
			resp, body := makeRequest(t, "GET", baseURL+"/posts/non-existent")
			
			assert.Equal(t, http.StatusNotFound, resp.StatusCode)
			
			var errorResponse map[string]interface{}
			err := json.Unmarshal(body, &errorResponse)
			require.NoError(t, err, "Failed to parse error response")
			
			assert.Contains(t, errorResponse, "error")
			errorMsg, ok := errorResponse["error"].(string)
			require.True(t, ok, "error should be a string")
			assert.Contains(t, errorMsg, "Post not found")
		})
	})

	t.Run("RSS endpoint", func(t *testing.T) {
		resp, body := makeRequest(t, "GET", baseURL+"/rss")
		
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		
		// Check content type
		contentType := resp.Header.Get("Content-Type")
		assert.Equal(t, "application/rss+xml; charset=utf-8", contentType)
		
		// Verify RSS structure
		rssContent := string(body)
		assert.Contains(t, rssContent, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>")
		assert.Contains(t, rssContent, "<rss version=\"2.0\">")
		assert.Contains(t, rssContent, "<channel>")
		assert.Contains(t, rssContent, "<title>My Blog</title>")
		assert.Contains(t, rssContent, "<description>Latest posts from my blog</description>")
		assert.Contains(t, rssContent, "</channel>")
		assert.Contains(t, rssContent, "</rss>")
		
		// Should contain our hello-world post
		assert.Contains(t, rssContent, "Hello World")
	})

	t.Run("Service health check", func(t *testing.T) {
		// Test that the service health check endpoint works
		// This verifies the service is responding correctly
		resp, body := makeRequest(t, "GET", baseURL+"/health")
		
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		
		var health map[string]interface{}
		err := json.Unmarshal(body, &health)
		require.NoError(t, err)
		
		assert.Equal(t, "healthy", health["status"])
		
		// Verify uptime is reasonable (should be just a few seconds since container started)
		uptime, ok := health["uptime"].(string)
		require.True(t, ok, "uptime should be a string")
		assert.NotEmpty(t, uptime)
	})

	t.Run("Invalid endpoints return 404", func(t *testing.T) {
		invalidEndpoints := []string{
			"/invalid",
			"/posts/",  // trailing slash
			"/api/posts", // wrong path
			"/health/invalid",
		}
		
		for _, endpoint := range invalidEndpoints {
			t.Run(fmt.Sprintf("GET %s", endpoint), func(t *testing.T) {
				resp, _ := makeRequest(t, "GET", baseURL+endpoint)
				assert.Equal(t, http.StatusNotFound, resp.StatusCode, 
					"Endpoint %s should return 404", endpoint)
			})
		}
	})

	t.Run("CORS headers", func(t *testing.T) {
		// Test that CORS headers are present
		client := &http.Client{Timeout: 10 * time.Second}
		
		req, err := http.NewRequest("OPTIONS", baseURL+"/posts", nil)
		require.NoError(t, err)
		
		// Add Origin header to trigger CORS
		req.Header.Set("Origin", "http://localhost:3000")
		req.Header.Set("Access-Control-Request-Method", "GET")
		
		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()
		
		// Should have CORS headers (depends on your CORS middleware implementation)
		// This test might need adjustment based on your specific CORS setup
		assert.Contains(t, resp.Header, "Access-Control-Allow-Origin")
	})
}

// TestServiceLogs demonstrates basic service testing (logs would be checked externally)
func TestServiceLogs(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping service log test in short mode")
	}
	
	baseURL := setupTest(t)
	
	// Make a request to generate some log activity
	resp, _ := makeRequest(t, "GET", baseURL+"/health")
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	
	t.Log("Service is responding correctly - logs would be checked externally via docker logs or other mechanisms")
}

// TestServicePerformance tests basic performance characteristics
func TestServicePerformance(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping performance test in short mode")
	}
	
	baseURL := setupTest(t)

	// Test response times
	endpoints := []string{
		"/health",
		"/posts", 
		"/posts/hello-world",
	}
	
	for _, endpoint := range endpoints {
		t.Run(fmt.Sprintf("Response time for %s", endpoint), func(t *testing.T) {
			start := time.Now()
			resp, _ := makeRequest(t, "GET", baseURL+endpoint)
			duration := time.Since(start)
			
			assert.Equal(t, http.StatusOK, resp.StatusCode)
			assert.Less(t, duration, 1*time.Second, 
				"Endpoint %s should respond within 1 second, took %v", endpoint, duration)
			
			t.Logf("Endpoint %s responded in %v", endpoint, duration)
		})
	}
}

// Benchmark for load testing (run with go test -bench=.)
func BenchmarkHealthEndpoint(b *testing.B) {
	baseURL := getBaseURL()
	
	// Quick check that service is available
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(baseURL + "/health")
	if err != nil {
		b.Fatalf("Service is not running at %s. Please start the service first. Error: %v", baseURL, err)
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b.Fatalf("Service health check failed. Expected status 200, got %d", resp.StatusCode)
	}
	
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			resp, err := client.Get(baseURL + "/health")
			if err != nil {
				b.Errorf("Request failed: %v", err)
				continue
			}
			if resp.StatusCode != http.StatusOK {
				b.Errorf("Expected status 200, got %d", resp.StatusCode)
			}
			resp.Body.Close()
		}
	})
}