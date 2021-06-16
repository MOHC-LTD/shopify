package shopify

import "time"

// Products is a collection of products
type Products []Product

// Product is an item that can be bought by a customer.
type Product struct {
	// ID is the identifier of the product.
	ID int64
	// CreatedAt is the date and time the product was created.
	CreatedAt time.Time
	// BodyHTML is a description of the product. Supports HTML formatting.
	BodyHTML string
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
	Tags string
	// Title is the name of the product.
	Title string
	// UpdatedAt is the date and time the product was last updated.
	UpdatedAt time.Time
	// Variants is an array of product variants, each representing a different version of the product.
	Variants Variants
	// Vendor is the name of the products vendor
	Vendor string
}

// ProductRepository maintains the products of a shop.
type ProductRepository interface {
	// GetAll gets all products from Shopify
	GetAll() (Products, error)
}
