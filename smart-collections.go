package shopify

import "time"

// SmartCollections is a list of SmartCollections.
type SmartCollections []SmartCollection

// Smartcollection is a grouping of products defined by rules that are set by the merchant.
type SmartCollection struct {
	// BodyHTML is the description of the smart collection. Includes HTML markup. Many shop themes display this on the smart collection page.
	BodyHTML string
	// Handle is a human-friendly unique string for the smart collection. Automatically generated from the title. Used in shop themes by the Liquid templating language to refer to the smart collection. (maximum: 255 characters) */
	Handle string
	// ID is the ID for the smart collection.
	ID int64
	// Image is the image associated with the smart collection.
	Image Image
	// PublishedAt is the time and date when the smart collection was made visible. Returns 0 for a hidden collection.
	PublishedAt time.Time
	// PublishedScope is whether the smart collection is published to the Point of Sale channel.
	/*
		Valid Values:
		- web: The collection is published to the Online Store channel but not published to the Point of Sale channel.
		- global: The collection is published to both the Online Store channel and the Point of Sale channel.
	*/
	PublishedScope string
	// Rules is the list of rules that define what products go into the smart collection.
	Rules Rules
	// Disjunctive is whether the product must match all the rules to be included in the smart collection.
	/*
		Valid values:
			- true: Products only need to match one or more of the rules to be included in the smart collection.
			- false: Products must match all of the rules to be included in the smart collection.
	*/
	Disjunctive bool
	// SortOrder is the order in which products in the collection appear.
	/*
		Valid values:
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
	// TemplateSuffix is the suffix of the liquid template being used.
	/*
		For example, if the value is custom, then the collection is using the collection.custom.liquid template.
		If the value is null, then the collection is using the default collection.liquid.
	*/
	TemplateSuffix string
	// Title is the name of the smart collection. (limit: 255 characters)
	Title string
	// ProductCount is the number of products that are in the collection
	ProductCount int
	// UpdatedAt is the date and time when the smart collection was last modified.
	UpdatedAt time.Time
}

// SmartCollectionRepository maintains the collections of a shop.
type SmartCollectionRepository interface {
	// Get retrieves a single smart collection
	Get(id int64) (SmartCollection, error)
}
