package content

// ContentDir is the root directory for content files
const ContentDir = "content"

// type Post will load a folder called posts
type Post struct {
	// Frontmatter
	Title       string `yaml:"title"`
	Date        string `yaml:"date"`
	Description string `yaml:"description"`
}
