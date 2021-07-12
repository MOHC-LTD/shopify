package shopify

import "time"

// Collections is a list of collections
type Collections []Collection

// Collection is a grouping of products.
type Collection struct {
	// ID is the ID for the collection.
	ID int64
	// Title is the name of the collection. (limit: 255 characters)
	Title string
	// BodyHTML is a description of the collection, complete with HTML markup. Many templates display this on their collection pages.
	BodyHTML string
	// Handle is a unique, human-readable string for the collection automatically generated from its title.
	/* This is used in themes by the Liquid templating language to refer to the collection. (limit: 255 characters) */
	Handle string
	// Image is the image assoicated with collection
	Image Images
	// PublishedAt is the time and date (ISO 8601 format) when the collection was made visible. Returns null for a hidden collection.
	PublishedAt time.Time
	// PublishedScope is whether the collection is published to the Point of Sale channel.
	/*
		Valid Values:
		- web: The collection is published to the Online Store channel but not published to the Point of Sale channel.
		- global: The collection is published to both the Online Store channel and the Point of Sale channel.
	*/
	PublishedScope string
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
	SortOrder string
	// TemplateSuffix is The suffix of the liquid template being used.
	/*
		For example, if the value is custom, then the collection is using the collection.custom.liquid template.
		If the value is null, then the collection is using the default collection.liquid.
	*/
	TemplateSuffix string
	//UpdatedAt is the date and time (ISO 8601 format) when the collection was last modified.
	UpdatedAt time.Time
}

// CollectionQuery are properties that can be used to filter the returned collections
// See https://shopify.dev/api/admin/rest/reference/products/collection
type CollectionQuery struct {
	/*
		Return only products specified by a list of collection IDs.
	*/
	IDs []int64
}

// CollectionRepository maintains the collections of a shop.
type CollectionRepository interface {
	// List gets all of the Collections
	List(query ProductQuery) (Collections, error)
}
