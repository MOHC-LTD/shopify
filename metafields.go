package shopify

import (
	"strconv"
	"time"
)

// Metafields are a flexible way to attach additional information to a Shopify resource (e.g. Product, Collection, etc.).
type Metafields []Metafield

// GetByKey gets a metafield via its key
func (m Metafields) GetByKey(key string) (interface{}, error) {
	for _, metafield := range m {
		if metafield.Key == key {
			switch metafield.Type {
			case NumberIntegerMetaFieldType:
				var value int64
				if metafield.Value != "" {
					return strconv.ParseInt(metafield.Value, 0, 64)
				}

				return value, nil
			case BooleanMetaFieldType:
				return strconv.ParseBool(metafield.Value)
			default:
				return metafield.Value, nil
			}
		}
	}

	return "", nil
}

// OwnerResource is the type of resource that the metafield is attached to.
type OwnerResource string

// MetafieldType is the type of data that the metafield stores in the `value` field.
type MetafieldType string

// MetafieldListType is the type of data that the metafield stores in the `value` field for list types.
type MetafieldListType MetafieldType

// Types
const (
	// BooleanMetaFieldType is a true or false value.
	BooleanMetaFieldType MetafieldType = "boolean"
	// ColorMetaFieldType is a hexadecimal code for a color.
	ColorMetaFieldType MetafieldType = "color"
	// DateMetaFieldType is a date in ISO 8601 format without a presumed timezone.
	DateMetaFieldType MetafieldType = "date"
	// DateTimeMetaFieldType is a date and time in ISO 8601 format without a presumed timezone
	DateTimeMetaFieldType MetafieldType = "date_time"
	// DimensionMetaFieldType is a value and a unit of length. Valid unit values: in, ft, yd, mm, cm, m
	DimensionMetaFieldType MetafieldType = "dimension"
	// FileReferenceMetaFieldType is a reference to a file on the online store. The default value is GenericFile. You can use validations to add other file types (for example, Image).
	FileReferenceMetaFieldType MetafieldType = "file_reference"
	// JSONMetaFieldType is a JSON-serializable value. This can be an object, an array, a string, a number, a boolean, or a null value.
	JSONMetaFieldType MetafieldType = "json"
	// MultiLineTextFieldMetaFieldType is a multi-line text field.
	MultiLineTextFieldMetaFieldType MetafieldType = "multi_line_text_field"
	// NumberDecimalMetaFieldType is a number with decimal places in the range of +/-9999999999999.999999999.
	NumberDecimalMetaFieldType MetafieldType = "number_decimal"
	// NumberIntegerMetaFieldType is a whole number in the range of +/-9,007,199,254,740,991.
	NumberIntegerMetaFieldType MetafieldType = "number_integer"
	// PageReferenceMetaFieldType is a reference to a page on the online store.
	PageReferenceMetaFieldType MetafieldType = "page_reference"
	// ProductReferenceMetaFieldType is a reference to a product on the online store.
	ProductReferenceMetaFieldType MetafieldType = "product_reference"
	// RatingMetaFieldType is a rating measured on a specified scale. Validations are required for ratings.
	RatingMetaFieldType MetafieldType = "rating"
	// SingleLineTextFieldMetaFieldType is a single-line text field.
	SingleLineTextFieldMetaFieldType MetafieldType = "single_line_text_field"
	// URLMetaFieldType is a URL with one of the allowed schemes: https, http, mailto, sms, tel.
	URLMetaFieldType MetafieldType = "url"
	// VariantReferenceMetaFieldType is a reference to a product variant on the online store.
	VariantReferenceMetaFieldType MetafieldType = "variant_reference"
	// VolumeMetaFieldType is a value and a unit of volume. Valid unit values: ml, cl, l, m3 (cubic meters), us_fl_oz, us_pt, us_qt, us_gal, imp_fl_oz, imp_pt, imp_qt, imp_gal.
	VolumeMetaFieldType MetafieldType = "volume"
	// WeightMetaFieldType is a value and a unit of weight. Valid unit values: oz, lb, g, kg.
	WeightMetaFieldType MetafieldType = "weight"
)

// List types
const (
	// ListColorMetaFieldType is a list of hexadecimal color codes.
	ListColorMetaFieldType MetafieldListType = "list.color"
	// ListDateMetaFieldType is a list of dates in ISO 8601 format without presumed timezones.
	ListDateMetaFieldType MetafieldListType = "list.date"
	// ListDateTimeMetaFieldType is a list of dates and times in ISO 8601 format without presumed timezones.
	ListDateTimeMetaFieldType MetafieldListType = "list.date_time"
	// ListDimensionMetaFieldType is a list of values and a unit of length. Valid unit values: in, ft, yd, mm, cm, m.
	ListDimensionMetaFieldType MetafieldListType = "list.date_time"
	// ListFileReferenceMetaFieldType is a list references to a file on the online store. The default value is GenericFile. You can use validations to add other file types (for example, Image
	ListFileReferenceMetaFieldType MetafieldListType = "list.file_reference"
	// ListNumberIntegerMetaFieldType is a list of whole numbers in the range of +/-9,007,199,254,740,991.
	ListNumberIntegerMetaFieldType MetafieldListType = "list.number_integer"
	// ListNumberDecimalMetaFieldType is a list of numbers with decimal places in the range of +/-9999999999999.999999999.
	ListNumberDecimalMetaFieldType MetafieldListType = "list.number_decimal"
	// ListPageReferenceMetaFieldType is a list of references to pages on the online store.
	ListPageReferenceMetaFieldType MetafieldListType = "list.page_reference"
	// ListProductReferenceMetaFieldType is a list of product references.
	ListProductReferenceMetaFieldType MetafieldListType = "list.product_reference"
	// ListRatingMetaFieldType is a list of ratings measured on a specified scale. Validations are required for ratings.
	ListRatingMetaFieldType MetafieldListType = "list.product_reference"
	// ListSingleLineTextFieldMetaFieldType is a list of single-line text fields.
	ListSingleLineTextFieldMetaFieldType MetafieldListType = "list.single_line_text_field"
	// ListURLMetaFieldType is a list of URLs with one of the allowed schemes: https, http, mailto, sms, tel.
	ListURLMetaFieldType MetafieldListType = "list.url"
	// ListVariantReferenceMetaFieldType is a list of references to a product variant on the online store.
	ListVariantReferenceMetaFieldType MetafieldListType = "list.variant_reference"
	// ListVolumeMetaFieldType is a list of values and a unit of volume. Valid unit values: ml, cl, l, m3 (cubic meters), us_fl_oz, us_pt, us_qt, us_gal, imp_fl_oz, imp_pt, imp_qt, imp_gal.
	ListVolumeMetaFieldType MetafieldListType = "list.volume"
	// ListWeightMetaFieldType is a list of values and a unit of weight. Valid unit values: oz, lb, g, kg
	ListWeightMetaFieldType MetafieldListType = "list.weight"
)

// Resources
const (
	// ArticleResource is the resource for metafields related to articles.
	ArticleResource OwnerResource = "article"
	// BlogResource is the resource for metafields related to the blog.
	BlogResource OwnerResource = "blog"
	// CollectionResource is the resource for metafields related to collections.
	CollectionResource OwnerResource = "collection"
	// CustomerResource is the resource for metafields related to customers.
	CustomerResource OwnerResource = "customer"
	// DraftOrderResource is the resource for metafields related to draft orders.
	DraftOrderResource OwnerResource = "draft_order"
	// LocationResource is the resource for metafields related to locations.
	LocationResource OwnerResource = "location"
	// OrderResource is the resource for metafields related to orders.
	OrderResource OwnerResource = "order"
	// PageResource is the resource for metafields related to pages.
	PageResource OwnerResource = "page"
	// ProductResource is the resource for metafields related to products.
	ProductResource OwnerResource = "product"
	// ProductImageResource is the resource for metafields related to product images.
	ProductImageResource OwnerResource = "product_image"
	// ProductVariantResource is the resource for metafields related to product variants.
	ProductVariantResource OwnerResource = "variants"
	// ShopResource is the resource for metafields related to the shop.
	ShopResource OwnerResource = "shop"
)

// Metafield is a flexible way to attach additional information to a Shopify resource (e.g. Product, Collection, etc.).
type Metafield struct {
	// ID is the unique ID of the metafield.
	ID int64
	// Description is the description of the information that the metafield contains.
	Description string
	// Key is the key of the metafield. Keys can be up to 30 characters long and can contain alphanumeric characters, hyphens, underscores, and periods.
	Key string
	// Namespace is the container name for a group of metafields. Grouping metafields within a namespace prevents your metafields from conflicting with other metafields with the same key name. Must have between 3-20 characters.
	Namespace string
	// Resource is the unique ID of the resource that the metafield is attached to and the type of resource that the metafield is attached to.
	Resource MetafieldResource
	// Value is the data stored in the metafield. The value is always stored as a string, regardless of the metafield's type.
	Value string
	// MetafieldType is the type of data that the metafield stores in the `value` field.
	Type MetafieldType
	// CreatedAt is the  date and time (ISO 8601 format) when the metafield was created.
	CreatedAt time.Time
	// UpdatedAt is the  date and time (ISO 8601 format) when the metafield was last updated.
	UpdatedAt time.Time
}

// MetafieldResource holds the query data for a resource
type MetafieldResource struct {
	// OwnerID is the unique ID of the resource that the metafield is attached to.
	OwnerID int64
	// OwnerResource is the type of resource that the metafield is attached to.
	OwnerResource OwnerResource
}

// MetaFieldRepository maintains the metafields of a shop.
type MetaFieldRepository interface {
	// List gets all the metafields
	List(query MetafieldQuery) (Metafields, error)
}

// MetafieldQuery are properties that can be used to filter the returned metafields
// See https://shopify.dev/api/admin-rest/2022-07/resources/metafield#get-metafields
type MetafieldQuery struct {
	// Resource is the resource and ID that the metafields are attached to
	Resource MetafieldResource
}
