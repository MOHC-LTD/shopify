package shopify

import "time"

// CustomCollections is a list of CustomCollections.
type CustomCollections []CustomCollection

// CustomCollection is a grouping of products that a merchant can create to make their store easier to browse.
type CustomCollection struct {
	bodyHTML       string
	collectionType string
	handle         string
	id             int64
	image          Image
	productCount   int
	publishedAt    time.Time
	publishedScope string
	sortOrder      string
	templateSuffix string
	title          string
	updatedAt      time.Time
}

// BodyHTML is the description of the custom collection. Includes HTML markup. Many shop themes display this on the custom collection page.
func (customCollection CustomCollection) BodyHTML() string {
	return customCollection.bodyHTML
}

// Handle is a human-friendly unique string for the custom collection. Automatically generated from the title. Used in shop themes by the Liquid templating language to refer to the custom collection. (maximum: 255 characters) */
func (customCollection CustomCollection) Handle() string {
	return customCollection.handle
}

// ID is the ID for the custom collection.
func (customCollection CustomCollection) ID() int64 {
	return customCollection.id
}

// Image is the image associated with the custom collection.
func (customCollection CustomCollection) Image() Image {
	return customCollection.image
}

// ProductCount is the number of products that are in the collection
func (customCollection CustomCollection) ProductCount() int {
	return customCollection.productCount
}

// PublishedAt is the time and date when the custom collection was made visible. Returns 0 for a hidden collection.
func (customCollection CustomCollection) PublishedAt() time.Time {
	return customCollection.publishedAt
}

// PublishedScope is whether the custom collection is published to the Point of Sale channel.
/*
	Valid Values:
	- web: The collection is published to the Online Store channel but not published to the Point of Sale channel.
	- global: The collection is published to both the Online Store channel and the Point of Sale channel.
*/
func (customCollection CustomCollection) PublishedScope() string {
	return customCollection.publishedScope
}

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
func (customCollection CustomCollection) SortOrder() string {
	return customCollection.sortOrder
}

// TemplateSuffix is the suffix of the liquid template being used.
/*
	For example, if the value is custom, then the collection is using the collection.custom.liquid template.
	If the value is null, then the collection is using the default collection.liquid.
*/
func (customCollection CustomCollection) TemplateSuffix() string {
	return customCollection.templateSuffix
}

// Title is the name of the custom collection. (limit: 255 characters)
func (customCollection CustomCollection) Title() string {
	return customCollection.title
}

// UpdatedAt is the date and time when the custom collection was last modified.
func (customCollection CustomCollection) UpdatedAt() time.Time {
	return customCollection.updatedAt
}

// NewCustomCollection create new CustomCollection
func NewCustomCollection(
	bodyHTML string,
	collectionType string,
	handle string,
	id int64,
	image Image,
	productCount int,
	publishedAt time.Time,
	publishedScope string,
	rules Rules,
	disjunctive bool,
	sortOrder string,
	templateSuffix string,
	title string,
	updatedAt time.Time,
) CustomCollection {
	return CustomCollection{
		bodyHTML:       bodyHTML,
		collectionType: collectionType,
		handle:         handle,
		id:             id,
		image:          image,
		productCount:   productCount,
		publishedAt:    publishedAt,
		publishedScope: publishedScope,
		sortOrder:      sortOrder,
		templateSuffix: templateSuffix,
		title:          title,
		updatedAt:      updatedAt,
	}
}

// IsCustomCollection checks if collection is smart collection
func (customCollection CustomCollection) IsSmartCollection() bool {
	return false
}

// IsCustomCollection checks if collection is custom collection
func (customCollection CustomCollection) IsCustomCollection() bool {
	return false
}
