package shopify

// ShippingLines is a collection of shipping lines
type ShippingLines []ShippingLine

// ShippingLine contains details about the shipping method used.
type ShippingLine struct {
	// Code is a reference to the shipping method.
	Code string
	// Price is the cost of the shipping method, in the store's currency
	Price string
	// PriceSet is the price of the shipping method in shop and presentment currencies.
	PriceSet PriceSet
	// DiscountedPrice is the price of the shipping method after line-level discounts have been applied. Doesn't reflect cart-level or order-level discounts.
	DiscountedPrice string
	// DiscountedPriceSet is the price of the shipping method in both shop and presentment currencies after line-level discounts have been applied.
	DiscountedPriceSet PriceSet
	// ID is the ID of the shipping method.
	ID int64
	// Title is the title of the shipping method
	Title string
}
