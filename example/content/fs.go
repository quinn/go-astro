// Code generated by go-astro. DO NOT EDIT.
package content

import (
	"embed"
	"fmt"

	"go.quinn.io/go-astro/content"
)

//go:embed posts
var FS embed.FS
type PostItem = content.ContentItem[Post]

// Initialize loads all content from the embedded filesystem.
// This must be called before using any Get* functions.
func Initialize() error {
	if err := content.LoadItems[Post](FS); err != nil {
		return fmt.Errorf("failed to load posts: %w", err)
	}
	return nil
}
// GetPosts returns all posts with their metadata and content.
func GetPosts() ([]PostItem, error) {
	return content.GetItems[Post]()
}
