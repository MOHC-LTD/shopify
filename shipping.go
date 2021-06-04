package shopify

// ShippingLine contains details about the shipping method used.
type ShippingLine struct {
	// ID is the ID of the shipping method.
	ID int64
	// Code is a reference to the shipping method.
	Code string
	// Title is the title of the shipping method
	Title string
}
