package shopify

// ShippingLine contains details about the shipping method used.
type ShippingLine struct {
	// ID is the ID of the shipping method.
	ID int64
	// Code is a reference to the shipping method.
	Code string
	// Title is the title of the shipping method
	Title string
	// Price is the cost of the shipping method, in the store's currency
	Price Money
	// Discount is the amount the shipping method is discounted by, in the store's currency
	Discount Money
}
