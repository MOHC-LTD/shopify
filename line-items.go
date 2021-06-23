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
	// Price is the cost of the item before any discounts or tax are applied, in the store's currency
	Price Money
	// Discount is the amount of money discounted from the price, in the store's currency
	Discount Money
	// Taxes is all the taxes applied to the item
	Taxes []Tax
}
