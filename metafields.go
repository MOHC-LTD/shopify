package shopify

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// ConvertToInt converts strings, float64 or int64 to int
func ConvertToInt(v any) (int, error) {
	switch value := any(v).(type) {
	case int:
		return value, nil
	case float64:
		return int(value), nil
	case int64:
		return int(value), nil
	case string:
		converted, err := strconv.ParseInt(v.(string), 0, 64)
		if err != nil {
			fmt.Println("Error parsing string to int:", err)
			return 0, err
		}
		return int(converted), nil
	default:
		return 0, fmt.Errorf("unsupported type %T", v)
	}
}

// Metafields are a flexible way to attach additional information to a Shopify resource (e.g. Product, Collection, etc.).
type Metafields []Metafield

// GetByKey gets a metafield via its key and namespace. Note defaults to string if type is not found. This is because most
// types are strings and there are a lot of different metafield types that would need to be caught.
func (m Metafields) GetByKey(key string, namespace string) (Metafield, error) {
	for _, metafield := range m {
		if metafield.Key == key && metafield.Namespace == namespace {
			switch metafield.Type {
			case NumberIntegerMetaFieldType:
				converted, err := ConvertToInt(metafield.Value)
				if err != nil {
					return Metafield{}, fmt.Errorf(
						"could not convert %v metafield type from key %v",
						NumberIntegerMetaFieldType,
						key,
					)
				}

				metafield.Value = converted

				return metafield, nil
			case BooleanMetaFieldType:
				converted, err := strconv.ParseBool(metafield.Value.(string))
				if err != nil {
					return Metafield{}, fmt.Errorf(
						"could not convert %v metafield type from key %v",
						NumberIntegerMetaFieldType,
						key,
					)
				}

				metafield.Value = converted

				return metafield, nil
			case ListSingleLineTextFieldMetaFieldType:
				var converted []string
				err := json.Unmarshal([]byte(metafield.Value.(string)), &converted)
				if err != nil {
					return Metafield{}, fmt.Errorf(
						"could not convert %v metafield type from key %v",
						ListSingleLineTextFieldMetaFieldType,
						key,
					)
				}

				metafield.Value = converted

				return metafield, nil
			default:
				return metafield, nil
			}
		}
	}

	return Metafield{}, NewErrMetafieldNotFoundByKey(key)
}

// GetByNamespace gets a list of metafields under the same namespace.
func (m Metafields) GetByNamespace(namespace string) (Metafields, error) {
	meta := make(Metafields, 0, len(m))

	for _, metafield := range m {
		if metafield.Namespace == namespace {
			meta = append(meta, metafield)
		}
	}

	return meta, nil
}

// Types
const (
	// BooleanMetaFieldType is a true or false value.
	BooleanMetaFieldType = "boolean"
	// ColorMetaFieldType is a hexadecimal code for a color.
	ColorMetaFieldType = "color"
	// DateMetaFieldType is a date in ISO 8601 format without a presumed timezone.
	DateMetaFieldType = "date"
	// DateTimeMetaFieldType is a date and time in ISO 8601 format without a presumed timezone
	DateTimeMetaFieldType = "date_time"
	// DimensionMetaFieldType is a value and a unit of length. Valid unit values: in, ft, yd, mm, cm, m
	DimensionMetaFieldType = "dimension"
	// FileReferenceMetaFieldType is a reference to a file on the online store. The default value is GenericFile. You can use validations to add other file types (for example, Image).
	FileReferenceMetaFieldType = "file_reference"
	// JSONMetaFieldType is a JSON-serializable value. This can be an object, an array, a string, a number, a boolean, or a null value.
	JSONMetaFieldType = "json"
	// MultiLineTextFieldMetaFieldType is a multi-line text field.
	MultiLineTextFieldMetaFieldType = "multi_line_text_field"
	// NumberDecimalMetaFieldType is a number with decimal places in the range of +/-9999999999999.999999999.
	NumberDecimalMetaFieldType = "number_decimal"
	// NumberIntegerMetaFieldType is a whole number in the range of +/-9,007,199,254,740,991.
	NumberIntegerMetaFieldType = "number_integer"
	// PageReferenceMetaFieldType is a reference to a page on the online store.
	PageReferenceMetaFieldType = "page_reference"
	// ProductReferenceMetaFieldType is a reference to a product on the online store.
	ProductReferenceMetaFieldType = "product_reference"
	// RatingMetaFieldType is a rating measured on a specified scale. Validations are required for ratings.
	RatingMetaFieldType = "rating"
	// SingleLineTextFieldMetaFieldType is a single-line text field.
	SingleLineTextFieldMetaFieldType = "single_line_text_field"
	// URLMetaFieldType is a URL with one of the allowed schemes: https, http, mailto, sms, tel.
	URLMetaFieldType = "url"
	// VariantReferenceMetaFieldType is a reference to a product variant on the online store.
	VariantReferenceMetaFieldType = "variant_reference"
	// VolumeMetaFieldType is a value and a unit of volume. Valid unit values: ml, cl, l, m3 (cubic meters), us_fl_oz, us_pt, us_qt, us_gal, imp_fl_oz, imp_pt, imp_qt, imp_gal.
	VolumeMetaFieldType = "volume"
	// WeightMetaFieldType is a value and a unit of weight. Valid unit values: oz, lb, g, kg.
	WeightMetaFieldType = "weight"
)

// List types
const (
	// ListColorMetaFieldType is a list of hexadecimal color codes.
	ListColorMetaFieldType = "list.color"
	// ListDateMetaFieldType is a list of dates in ISO 8601 format without presumed timezones.
	ListDateMetaFieldType = "list.date"
	// ListDateTimeMetaFieldType is a list of dates and times in ISO 8601 format without presumed timezones.
	ListDateTimeMetaFieldType = "list.date_time"
	// ListDimensionMetaFieldType is a list of values and a unit of length. Valid unit values: in, ft, yd, mm, cm, m.
	ListDimensionMetaFieldType = "list.date_time"
	// ListFileReferenceMetaFieldType is a list references to a file on the online store. The default value is GenericFile. You can use validations to add other file types (for example, Image
	ListFileReferenceMetaFieldType = "list.file_reference"
	// ListNumberIntegerMetaFieldType is a list of whole numbers in the range of +/-9,007,199,254,740,991.
	ListNumberIntegerMetaFieldType = "list.number_integer"
	// ListNumberDecimalMetaFieldType is a list of numbers with decimal places in the range of +/-9999999999999.999999999.
	ListNumberDecimalMetaFieldType = "list.number_decimal"
	// ListPageReferenceMetaFieldType is a list of references to pages on the online store.
	ListPageReferenceMetaFieldType = "list.page_reference"
	// ListProductReferenceMetaFieldType is a list of product references.
	ListProductReferenceMetaFieldType = "list.product_reference"
	// ListRatingMetaFieldType is a list of ratings measured on a specified scale. Validations are required for ratings.
	ListRatingMetaFieldType = "list.product_reference"
	// ListSingleLineTextFieldMetaFieldType is a list of single-line text fields.
	ListSingleLineTextFieldMetaFieldType = "list.single_line_text_field"
	// ListURLMetaFieldType is a list of URLs with one of the allowed schemes: https, http, mailto, sms, tel.
	ListURLMetaFieldType = "list.url"
	// ListVariantReferenceMetaFieldType is a list of references to a product variant on the online store.
	ListVariantReferenceMetaFieldType = "list.variant_reference"
	// ListVolumeMetaFieldType is a list of values and a unit of volume. Valid unit values: ml, cl, l, m3 (cubic meters), us_fl_oz, us_pt, us_qt, us_gal, imp_fl_oz, imp_pt, imp_qt, imp_gal.
	ListVolumeMetaFieldType = "list.volume"
	// ListWeightMetaFieldType is a list of values and a unit of weight. Valid unit values: oz, lb, g, kg
	ListWeightMetaFieldType = "list.weight"
)

// Resources
const (
	// ArticleResource is the resource for metafields related to articles.
	ArticleResource = "article"
	// BlogResource is the resource for metafields related to the blog.
	BlogResource = "blog"
	// CollectionResource is the resource for metafields related to collections.
	CollectionResource = "collection"
	// CustomerResource is the resource for metafields related to customers.
	CustomerResource = "customer"
	// DraftOrderResource is the resource for metafields related to draft orders.
	DraftOrderResource = "draft_order"
	// LocationResource is the resource for metafields related to locations.
	LocationResource = "location"
	// OrderResource is the resource for metafields related to orders.
	OrderResource = "order"
	// PageResource is the resource for metafields related to pages.
	PageResource = "page"
	// ProductResource is the resource for metafields related to products.
	ProductResource = "product"
	// ProductImageResource is the resource for metafields related to product images.
	ProductImageResource = "product_image"
	// ProductVariantResource is the resource for metafields related to product variants.
	ProductVariantResource = "variants"
	// ShopResource is the resource for metafields related to the shop.
	ShopResource = "shop"
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
	// Value is the data stored in the metafield. The model stores the metafield in it's actual type.
	Value interface{}
	// MetafieldType is the type of data that the metafield stores in the `value` field.
	Type string
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
	OwnerResource string
}

// MetafieldRepository maintains the metafields of a shop.
type MetafieldRepository interface {
	// List gets all the metafields
	List(query MetafieldQuery) (Metafields, error)
}

// MetafieldQuery are properties that can be used to filter the returned metafields
// See https://shopify.dev/api/admin-rest/2022-07/resources/metafield#get-metafields
type MetafieldQuery struct {
	// Resource is the resource and ID that the metafields are attached to
	Resource MetafieldResource
}

// ErrMetafieldNotFoundByKey is thrown when a metafield could not be found by its key
type ErrMetafieldNotFoundByKey struct {
	key string
}

func (err ErrMetafieldNotFoundByKey) Error() string {
	return fmt.Sprintf("could not find metafield with key %v", err.key)
}

// NewErrMetafieldNotFoundByKey builds the error
func NewErrMetafieldNotFoundByKey(key string) ErrMetafieldNotFoundByKey {
	return ErrMetafieldNotFoundByKey{key}
}
