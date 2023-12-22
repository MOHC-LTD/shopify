package shopify

import (
	"time"
)

// Articles is a list of Articles
type Articles []Article

// Article is a shopify article
type Article struct {
	// Author is the name of the author of the article.
	Author string
	// BlogID is the ID of the blog containing the article.
	BlogID string
	// BodyHTML is the text of the body of the article, complete with HTML markup.
	BodyHTML string
	// CreatedAt The date and time (ISO 8601 format) when the article was created.
	CreatedAt time.Time
	// ID is the ID of the article.
	ID int64
	// Handle is a unique, human-readable string for the article automatically generated from its title.
	Handle string
	// Image is the image associated with the article.
	Image Image
	// Published is whether the article is visible.
	Published bool
	// PublishedAt is the time and date when the article was published.
	PublishedAt time.Time
	// SummaryHTML is a summary of the article, which can include HTML markup. The summary is used by the
	// online store theme to display the article on other pages, such as the home page or the main blog page.
	SummaryHTML string
	// Tags is a comma-separated list of tags. Tags are additional short descriptors formatted as a string of comma-separated values.
	Tags string
	// Title is the name of the article. (limit: 255 characters)
	Title string
	// UpdatedAt is the date and time when the article was last modified.
	UpdatedAt time.Time
	// UserId is a unique numeric identifier for the author of the article.
	UserID int64
}

// ArticleRepository maintains the Articles of a shop.
type ArticleRepository interface {
	// Get retrieves a single article
	Get(blogID, id int64) (Article, error)
	// GetAll retrieves a list of all articles in a blog
	GetAll(blogID int64) (Articles, error)
}
