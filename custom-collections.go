package shopify

import "time"

// CustomCollections is a list of CustomCollections.
type CustomCollections []CustomCollection

// CustomCollection is a grouping of products that a merchant can create to make their store easier to browse.
type CustomCollection struct {
	bodyHTML       string
	handle         string
	id             int64
	image          Image
	publishedAt    time.Time
	publishedScope string
	sortOrder      string
	templateSuffix string
	title          string
	productCount   int
	updatedAt      time.Time
}

func (customCollection CustomCollection) BodyHTML() string {
	return customCollection.bodyHTML
}

func (customCollection CustomCollection) Handle() string {
	return customCollection.handle
}
func (customCollection CustomCollection) Image() Image {
	return customCollection.image
}

func (customCollection CustomCollection) ID() int64 {
	return customCollection.id
}

func (customCollection CustomCollection) ProductCount() int {
	return customCollection.productCount
}

func (customCollection CustomCollection) PublishedAt() time.Time {
	return customCollection.publishedAt
}

func (customCollection CustomCollection) PublishedScope() string {
	return customCollection.publishedScope
}

func (customCollection CustomCollection) SortOrder() string {
	return customCollection.sortOrder
}

func (customCollection CustomCollection) TemplateSuffix() string {
	return customCollection.templateSuffix
}

func (customCollection CustomCollection) Title() string {
	return customCollection.title
}

func (customCollection CustomCollection) UpdatedAt() time.Time {
	return customCollection.updatedAt
}

// NewCustomCollection create new CustomCollection
func NewCustomCollection(
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
) CustomCollection {
	return CustomCollection{
		bodyHTML:       bodyHTML,
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

// CustomCollectionRepository maintains the collections of a shop.
type CustomCollectionRepository interface {
	// Get retrieves a single custom collection
	Get(id int64) (CustomCollection, error)
}
