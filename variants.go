package shopify

import "time"

// Variants is a collection of variants
type Variants []Variant

// Variant is a specific version of a product
type Variant struct {
	// ID is the unique numeric identifier for the product variant.
	ID int64
	// SKU is a unique identifier for the product variant in the shop.
	SKU string
	// Title is the title of the product variant.
	Title string
	// InventoryItemID is the unique identifier for the inventory item.
	InventoryItemID int64
	// Price is the price of the product variant.
	Price string
	// ProductID is the unique numeric identifier for the product.
	ProductID int64
	// Barcode is the barcode, UPC, or ISBN number for the product.
	Barcode string
	// CreatedAt is the date and time when the product variant was created.
	CreatedAt time.Time
	// CreatedAt is the date and time when the product variant was last modified.
	UpdatedAt time.Time
}

// VariantRepository maintains the product variants of a shop.
type VariantRepository interface {
	// Get gets a variant by its id
	Get(id int64) (Variant, error)
}
