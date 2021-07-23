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

func (smartCollection SmartCollection) BodyHTML() string {
	return smartCollection.bodyHTML
}

func (smartCollection SmartCollection) Handle() string {
	return smartCollection.handle
}
func (smartCollection SmartCollection) Image() Image {
	return smartCollection.image
}

func (smartCollection SmartCollection) ID() int64 {
	return smartCollection.id
}

func (smartCollection SmartCollection) ProductCount() int {
	return smartCollection.productCount
}

func (smartCollection SmartCollection) PublishedAt() time.Time {
	return smartCollection.publishedAt
}

func (smartCollection SmartCollection) PublishedScope() string {
	return smartCollection.publishedScope
}

func (smartCollection SmartCollection) SortOrder() string {
	return smartCollection.sortOrder
}

func (smartCollection SmartCollection) TemplateSuffix() string {
	return smartCollection.templateSuffix
}

func (smartCollection SmartCollection) Title() string {
	return smartCollection.title
}

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
