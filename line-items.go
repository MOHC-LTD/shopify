package shopify

// LineItems is a collection of line items
type LineItems []LineItem

// LineItem is an item on an order
type LineItem struct {
	// ID is the ID of the line item
	ID int64
	// Title is the title of the product.
	Title string
	// Name is the name of the product variant.
	Name string
	// SKU is the item's SKU (stock keeping unit).
	SKU string
	// Quantity is the number of items that were purchased.
	Quantity int
}
