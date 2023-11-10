package shopify

import "time"

// Collections is a list of collections
type Collections []Collection

// Collection is a grouping of products.
type Collection interface {
	// BodyHTML is a description of the collection, complete with HTML markup. Many templates display this on their collection pages.
	BodyHTML() string
	// CollectionType is the type of collection
	CollectionType() string
	// Handle is a unique, human-readable string for the collection automatically generated from its title.
	/* This is used in themes by the Liquid templating language to refer to the collection. (limit: 255 characters) */
	Handle() string
	// ID is the ID for the collection.
	ID() int64
	// Image is the image associated with the collection.
	Image() Image
	// ProductsCount is the number of products that are in the collection
	ProductsCount() int
	// PublishedAt is the time and date when the collection was made visible. Returns 0 for a hidden collection.
	PublishedAt() time.Time
	// PublishedScope is whether the collection is published to the Point of Sale channel.
	/*
		Valid Values:
		- web: The collection is published to the Online Store channel but not published to the Point of Sale channel.
		- global: The collection is published to both the Online Store channel and the Point of Sale channel.
	*/
	PublishedScope() string
	// SortOrder is the order in which products in the collection appear.
	/*
		valid Values:
		- alpha-asc: Alphabetically, in ascending order (A - Z).
		- alpha-desc: Alphabetically, in descending order (Z - A).
		- best-selling: By best-selling products.
		- created: By date created, in ascending order (oldest - newest).
		- created-desc: By date created, in descending order (newest - oldest).
		- manual: In the order set manually by the shop owner.
		- price-asc: By price, in ascending order (lowest - highest).
		- price-desc: By price, in descending order (highest - lowest).
	*/
	SortOrder() string
	// TemplateSuffix is the suffix of the liquid template being used.
	/*
		For example, if the value is custom, then the collection is using the collection.custom.liquid template.
		If the value is null, then the collection is using the default collection.liquid.
	*/
	TemplateSuffix() string
	// Title is the name of the collection. (limit: 255 characters)
	Title() string
	// UpdatedAt is the date and time when the collection was last modified.
	UpdatedAt() time.Time
}

// CollectionRepository maintains the collections of a shop.
type CollectionRepository interface {
	// Get retrieves a single collection
	Get(id int64) (Collection, error)
	// GetAllSmartCollections retrieves all smart collections
	GetAllSmartCollections() (Collections, error)
	// GetAllCustomCollections retrieves all custom collections
	GetAllCustomCollections() (Collections, error)
	// Products retrieves a list product that belonging to a collection
	Products(id int64) (Products, error)
}
