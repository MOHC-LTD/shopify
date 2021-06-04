package shopify

import "time"

// FulfillmentEvent are events triggering during the fulfillment process
type FulfillmentEvent struct {
	// ID is the ID of the fulfillment event.
	ID int64
	// FulfillmentID is the ID for the fulfillment that's associated with the fulfillment event.
	FulfillmentID int64
	// Status is the status of the fulfillment event.
	/*
		Valid values:

		- label_printed: A label for the shipment was purchased and printed.
		- label_purchased: A label for the shipment was purchased, but not printed.
		- attempted_delivery: Delivery of the shipment was attempted, but unable to be completed.
		- ready_for_pickup: The shipment is ready for pickup at a shipping depot.
		- picked_up: The fulfillment was successfully picked up.
		- confirmed: The carrier is aware of the shipment, but hasn't received it yet.
		- in_transit: The shipment is being transported between shipping facilities on the way to its destination.
		- out_for_delivery: The shipment is being delivered to its final destination.
		- delivered: The shipment was successfully delivered.
		- failure: Something went wrong when pulling tracking information for the shipment, such as the tracking number was invalid or the shipment was canceled.

	*/
	Status string
	// CreatedAt is the date and time when the fulfillment event was created.
	CreatedAt time.Time
	// UpdateAt is the date and time when the fulfillment event was updated.
	UpdatedAt time.Time
}

// FulfillmentEventRepository manages fulfillment events
type FulfillmentEventRepository interface {
	// Create creates a new fulfillment event
	Create(orderID int64, fulfillmentID int64, event FulfillmentEvent) (FulfillmentEvent, error)
	// Delete deletes the specified fulfillment event
	Delete(orderID int64, fulfillmentID int64, eventID int64) error
	// List lists all of the fulfillment events for a specific fulfillment
	List(orderID int64, fulfillmentID int64) ([]FulfillmentEvent, error)
}
