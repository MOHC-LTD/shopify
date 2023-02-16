package shopify

import "time"

// FulfillmentOrders is a collection of fulfillment orders
type FulfillmentOrders []FulfillmentOrder

// FulfillmentOrder is a collection of line items for an order that can be
// fulfilled from the same location
type FulfillmentOrder struct {
	// ID is the ID of the fulfillment order
	ID int64
	// OrderID is the ID for the order that's associated with the fulfillment order
	OrderID int64
	// AssignedLocationID is the ID of the location assigned to do the work
	AssignedLocationID int64

	// Status is the status of the fulfillment order
	Status string
	// LineItems are the line items of the order to be fulfilled under this location
	LineItems FulfillmentOrderLineItems

	// CreatedAt is the date and time when the fulfillment order was created
	CreatedAt time.Time
	// UpdateAt is the date and time when the fulfillment order was updated
	UpdatedAt time.Time
}

// FulfillmentOrderLineItems is a collection of fulfillment order line items
type FulfillmentOrderLineItems []FulfillmentOrderLineItem

// FulfillmentOrderLineItem is a line item in a fulfillment order
type FulfillmentOrderLineItem struct {
	// ID is the ID of the fulfillment order line item
	ID int64
	// LineItemID is the ID of the line item in the order
	LineItemID int64
	// VariantID is the ID of the product variant
	VariantID int64
	// Quantity is the total number of units to be fulfilled
	Quantity int
	// FulfillableQuantity is the number of remaining units to be fulfilled
	FulfillableQuantity int
}

// FulfillmentOrderRepository manages fulfillment orders
type FulfillmentOrderRepository interface {
	// Get retrieves a single fulfillment order
	Get(id int64) (FulfillmentOrder, error)
	// List retrieves all fulfillment orders for a given order
	List(orderID int64) (FulfillmentOrders, error)
}
