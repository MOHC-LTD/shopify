package shopify

import "time"

// InventoryItemRepository manages inventory items
type InventoryItemRepository interface {
	// Get retrieves an inventory item by its id.
	Get(id int64) (InventoryItem, error)
}

// Country-specific Harmonized System (HS) codes for the item.
type CountryHarmonizedSystemCode struct {
	// The general Harmonized System (HS) code for the inventory item. Used if a country-specific HS code (`countryHarmonizedSystemCode`) is not available.
	HarmonizedSystemCode string
	// The country code
	CountryCode string
}

// InventoryLevel represents the amount of an item that there is available in a shop
type InventoryItem struct {
	// The ID of the inventory item
	ID int64
	// The unique SKU (stock keeping unit) of the inventory item
	SKU string
	// The date and time when the inventory level was created
	CreatedAt time.Time
	// The date and time when the inventory level was last modified
	UpdatedAt time.Time
	// Whether a customer needs to provide a shipping address when placing an order containing the inventory item
	RequiresShipping bool
	// The unit cost of the inventory item
	Cost string
	// The country code of where the item came from (ISO)
	CountryCodeOfOrigin string
	// The province code of where the item came from (ISO)
	ProvinceCodeOfOrigin string
	// The general Harmonized System (HS) code for the inventory item. Used if a country-specific HS code (`countryHarmonizedSystemCode`) is not available.
	HarmonizedSystemCode int64
	// Whether inventory levels are tracked for the item. If true, then the inventory quantity changes are tracked by Shopify.
	Tracked bool
	// An array of country-specific Harmonized System (HS) codes for the item. Used to determine duties when shipping the inventory item to certain countries.
	CountryHarmonizedSystemCodes []CountryHarmonizedSystemCode
	// The Graphql API for the inventory item
	AdminGraphqlApiId string
}
