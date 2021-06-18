package shopify

import "time"

// InventoryLevelRepository manages inventory levels
type InventoryLevelRepository interface {
	// Set sets the inventory level for an inventory item at a location
	Set(inventoryItemID int64, locationID int64, quantity int) (InventoryLevel, error)
}

// InventoryLevel represents the amount of an item that there is available in a shop
type InventoryLevel struct {
	// InventoryItemID is the ID of the inventory item that the inventory level belongs to
	InventoryItemID int64
	// Available is the quantity of inventory items available for sale
	Available int
	// LocationID is the ID of the location that the inventory level belongs to
	LocationID int64
	// The date and time when the inventory level was last modified
	UpdatedAt time.Time
}
