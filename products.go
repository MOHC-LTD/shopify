package shopify

import (
	"context"
	"time"
)

// Products is a collection of products
type Products []Product

// Product is an item that can be bought by a customer.
type Product struct {
	// BodyHTML is a description of the product. Supports HTML formatting.
	BodyHTML string
	// CreatedAt is the date and time the product was created.
	CreatedAt time.Time
	// Handle is a unique human-friendly string for the product.
	Handle string
	// ID is the identifier of the product.
	ID int64
	// Images are the images of the product
	Images ProductImages
	// Meta are the metafields of the product
	Meta []Metafield
	// ProductType is a categorization for the product used for filtering and searching products.
	ProductType string
	// PublishedAt is the date and time the product was published.
	PublishedAt time.Time
	// Status is the status of a product
	/*
		Valid values:
		active: The product is ready to sell and is available to customers on the online store, sales channels, and apps. By default, existing products are set to active.
		archived: The product is no longer being sold and isn't available to customers on sales channels and apps.
		draft: The product isn't ready to sell and is unavailable to customers on sales channels and apps. By default, duplicated and unarchived products are set to draft.
	*/
	Status string
	// Tags is a string of comma-separated tags that are used for filtering and search.
	/*
		A product can have up to 250 tags. Each tag can have up to 255 characters.
	*/
	Tags Tags
	// Title is the name of the product.
	Title string
	// UpdatedAt is the date and time the product was last updated.
	UpdatedAt time.Time
	// Variants is an array of product variants, each representing a different version of the product.
	Variants Variants
	// Vendor is the name of the products vendor
	Vendor string
	// Options are the options of a product
	/*
		The custom product properties. For example, Size, Color, and Material.
		Each product can have up to 3 options and each option value can be up to 255 characters.
		Product variants are made of up combinations of option values. Options cannot be created without
		values. To create new options, a variant with an associated option value also needs to be created.
	*/
	Options ProductOptions
}

// ProductOptions are the options of a product
type ProductOptions []ProductOption

// ProductOption is the option of a product
type ProductOption struct {
	// ID is an unsigned 64-bit integer that's used as a unique identifier for the product option.
	ID int64
	// Name is the custom name given to one of the three available product options that can be used.
	Name string
	// Position is the order of the product variants in the list.
	Position int
	// Values are the possible values that can be selected for the option and coincide with the product variant title, option1, option 2 or option3.
	Values []string
}

// ProductQuery are properties that can be used to filter the returned products
// See https://shopify.dev/docs/admin-api/rest/reference/products/product#index-2021-04
type ProductQuery struct {
	/*
		Return only products specified by a list of product IDs.
	*/
	IDs []int64
	/*
		Return only products specified by the tags provided
	*/
	Tags []string
}

// ProductRepository maintains the products of a shop.
type ProductRepository interface {
	// List gets all of the products
	List(query ProductQuery) (Products, error)
	// Get retrieves a single product
	Get(id int64) (Product, error)
	// Create creates a single product
	Create(product Product) (Product, error)
	// Update updates a single product
	Update(product Product) (Product, error)
	// Delete a single product
	Delete(productID int64) error
}

// ProductCreator supplies methods for creating products in the shopify shop
type ProductCreator interface {
	// Save saves a single product to the shopify shop and returns the created product
	Save(context.Context, Product) (Product, error)
}
