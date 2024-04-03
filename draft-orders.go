package shopify

// Orders is a collection of draft orders
type DraftOrders []DraftOrder

// DraftOrder is a customer's draft order.
type DraftOrder struct {
	// ID is the ID of the draft order
	ID int64
	// ShippingAddress is the mailing address to where the order will be shipped.
	ShippingAddress Address
}

// DraftOrderRepository maintains the draft orders in the shop
type DraftOrderRepository interface {
	// Get gets a draft order
	Get(id int64) (DraftOrder, error)
}
