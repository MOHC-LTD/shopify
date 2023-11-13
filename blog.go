package shopify

import (
	"time"
)

// Blogs is a list of blog
type Blogs []Blog

// Blogs is a shopify blog
type Blog struct {
	// Commentable indicates whether readers can post comments to the blog and if comments are moderated or no
	Commentable string
	// CreatedAt The date and time (ISO 8601 format) when the blog was created.
	CreatedAt time.Time
	// Handle is a unique, human-readable string for the blog automatically generated from its title.
	Handle string
	// ID is the ID of the blog.
	ID int64
	// Tags is a comma-separated list of tags. Tags are additional short descriptors formatted as a string of comma-separated values.
	Tags string
	// Title is the name of the blog. (limit: 255 characters)
	Title string
	// UpdatedAt is the date and time when the blog was last modified.
	UpdatedAt time.Time
}

// BlogRepository maintains the Blogs of a shop.
type BlogRepository interface {
	// Get retrieves a single blog
	Get(id int64) (Blog, error)
	// GetAll retrieves a list of all blogs
	GetAll() (Blogs, error)
}
