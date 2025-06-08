package models

import (
	"fmt"
	"html"
	"strings"
	"time"
)

// RSSFeed represents an RSS feed
type RSSFeed struct {
	Title       string
	Link        string
	Description string
	Language    string
	Items       []RSSItem
}

// RSSItem represents an RSS feed item
type RSSItem struct {
	Title       string
	Link        string
	Description string
	PubDate     string
	GUID        string
}

// ToXML converts the RSS feed to XML format
func (f *RSSFeed) ToXML() string {
	var items strings.Builder
	for _, item := range f.Items {
		items.WriteString(fmt.Sprintf(`
		<item>
			<title>%s</title>
			<link>%s</link>
			<description><![CDATA[%s]]></description>
			<pubDate>%s</pubDate>
			<guid>%s</guid>
		</item>`,
			html.EscapeString(item.Title),
			html.EscapeString(item.Link),
			item.Description,
			item.PubDate,
			html.EscapeString(item.GUID),
		))
	}

	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0">
	<channel>
		<title>%s</title>
		<link>%s</link>
		<description>%s</description>
		<language>%s</language>
		<lastBuildDate>%s</lastBuildDate>
		<generator>Blog API</generator>%s
	</channel>
</rss>`,
		html.EscapeString(f.Title),
		html.EscapeString(f.Link),
		html.EscapeString(f.Description),
		html.EscapeString(f.Language),
		time.Now().Format(time.RFC1123Z),
		items.String(),
	)
}

// BlogPostToRSSItem converts a BlogPost to an RSSItem
func BlogPostToRSSItem(post BlogPost, baseURL string) RSSItem {
	link := fmt.Sprintf("%s/posts/%s", strings.TrimRight(baseURL, "/"), post.Slug)

	// Use excerpt if available, otherwise truncate content
	description := post.Excerpt
	if description == "" && post.Content != "" {
		if len(post.Content) > 200 {
			description = post.Content[:200] + "..."
		} else {
			description = post.Content
		}
	}

	return RSSItem{
		Title:       post.Title,
		Link:        link,
		Description: description,
		PubDate:     time.Time(post.Date).Format(time.RFC1123Z),
		GUID:        link,
	}
}
