package shopify

import "time"

// SmartCollections is a list of SmartCollections.
type SmartCollections []SmartCollection

// SmartCollection is a grouping of products defined by rules that are set by the merchant.
type SmartCollection struct {
	bodyHTML       string
	handle         string
	id             int64
	image          Image
	productCount   int
	publishedAt    time.Time
	publishedScope string
	// Rules is the list of rules that define what products go into the smart collection.
	Rules Rules
	// Disjunctive is whether the product must match all the rules to be included in the smart collection.
	/*
		Valid values:
			- true: Products only need to match one or more of the rules to be included in the smart collection.
			- false: Products must match all of the rules to be included in the smart collection.
	*/
	Disjunctive    bool
	sortOrder      string
	templateSuffix string
	title          string
	updatedAt      time.Time
}

// BodyHTML is the description of the smart collection. Includes HTML markup. Many shop themes display this on the smart collection page.
func (smartCollection SmartCollection) BodyHTML() string {
	return smartCollection.bodyHTML
}

// Handle is a human-friendly unique string for the smart collection. Automatically generated from the title. Used in shop themes by the Liquid templating language to refer to the smart collection. (maximum: 255 characters) */
func (smartCollection SmartCollection) Handle() string {
	return smartCollection.handle
}

// ID is the ID for the smart collection.
func (smartCollection SmartCollection) ID() int64 {
	return smartCollection.id
}

// Image is the image associated with the smart collection.
func (smartCollection SmartCollection) Image() Image {
	return smartCollection.image
}

// ProductCount is the number of products that are in the collection
func (smartCollection SmartCollection) ProductCount() int {
	return smartCollection.productCount
}

// PublishedAt is the time and date when the smart collection was made visible. Returns 0 for a hidden collection.
func (smartCollection SmartCollection) PublishedAt() time.Time {
	return smartCollection.publishedAt
}

// PublishedScope is whether the smart collection is published to the Point of Sale channel.
/*
	Valid Values:
	- web: The collection is published to the Online Store channel but not published to the Point of Sale channel.
	- global: The collection is published to both the Online Store channel and the Point of Sale channel.
*/
func (smartCollection SmartCollection) PublishedScope() string {
	return smartCollection.publishedScope
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
func (smartCollection SmartCollection) SortOrder() string {
	return smartCollection.sortOrder
}

// TemplateSuffix is the suffix of the liquid template being used.
/*
	For example, if the value is custom, then the collection is using the collection.custom.liquid template.
	If the value is null, then the collection is using the default collection.liquid.
*/
func (smartCollection SmartCollection) TemplateSuffix() string {
	return smartCollection.templateSuffix
}

// Title is the name of the smart collection. (limit: 255 characters)
func (smartCollection SmartCollection) Title() string {
	return smartCollection.title
}

// UpdatedAt is the date and time when the smart collection was last modified.
func (smartCollection SmartCollection) UpdatedAt() time.Time {
	return smartCollection.updatedAt
}

// NewSmartCollection create new SmartColllection
func NewSmartCollection(
	bodyHTML string,
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
) SmartCollection {
	return SmartCollection{
		bodyHTML:       bodyHTML,
		handle:         handle,
		id:             id,
		image:          image,
		productCount:   productCount,
		publishedAt:    publishedAt,
		publishedScope: publishedScope,
		Rules:          rules,
		Disjunctive:    disjunctive,
		sortOrder:      sortOrder,
		templateSuffix: templateSuffix,
		title:          title,
		updatedAt:      updatedAt,
	}
}

// SmartCollectionRepository maintains the collections of a shop.
type SmartCollectionRepository interface {
	// Get retrieves a single smart collection
	Get(id int64) (SmartCollection, error)
}
