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
	// InventoryManagement is the fulfillment service that tracks the number of items in stock for the product variant.
	/*
		Valid values:

			- VariantInventoryManagementShopify: You are tracking inventory yourself using the admin.
			- VariantInventoryManagementNull: You aren't tracking inventory on the variant.
			- the handle of a fulfillment service that has inventory management enabled: This must be the same fulfillment service referenced by the FulfillmentService property.
	*/
	InventoryManagement string
	// InventoryQuantity is an aggregate of inventory across all locations. To adjust inventory at a specific location, use the InventoryLevel resource. Readonly.
	InventoryQuantity int
	// Price is the price of the product variant.
	Price string
	// CompareAtPrice is the price of the product variant before an adjustment or a sale.
	CompareAtPrice string
	// ProductID is the unique numeric identifier for the product.
	ProductID int64
	// Barcode is the barcode, UPC, or ISBN number for the product.
	Barcode string
	// CreatedAt is the date and time when the product variant was created.
	CreatedAt time.Time
	// CreatedAt is the date and time when the product variant was last modified.
	UpdatedAt time.Time
}

const (
	// VariantInventoryManagementShopify is the value of InventoryManagement when shopify is managing the inventory.
	VariantInventoryManagementShopify = "shopify"
	// VariantInventoryManagementNull is the value of InventoryManagement when the inventory is not being managed.
	VariantInventoryManagementNull = ""
)

// VariantRepository maintains the product variants of a shop.
type VariantRepository interface {
	// Get gets a variant by its id
	Get(id int64) (Variant, error)
}
