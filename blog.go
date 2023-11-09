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
	// CreatedAt The date and time (ISO 8601 format) when the article was created.
	CreatedAt time.Time
	// Handle is a unique, human-readable string for the collection automatically generated from its title.
	/* This is used in themes by the Liquid templating language to refer to the collection. (limit: 255 characters) */
	Handle string
	// ID is the ID for the collection.
	ID int64
	// Tags is a comma-separated list of tags. Tags are additional short descriptors formatted as a string of comma-separated values.
	Tags string
	// TemplateSuffix is the suffix of the liquid template being used.
	/*
		For example, if the value is custom, then the collection is using the collection.custom.liquid template.
		If the value is null, then the collection is using the default collection.liquid.
	*/
	TemplateSuffix string
	// Title is the name of the article. (limit: 255 characters)
	Title string
	// UpdatedAt is the date and time when the collection was last modified.
	UpdatedAt time.Time
}

// BlogRepository maintains the Blogs of a shop.
type BlogRepository interface {
	// Get retrieves a single blog
	Get(id int64) (Blog, error)
	// GetAll retrieves a list of all blogs
	GetAll() (Blogs, error)
}
