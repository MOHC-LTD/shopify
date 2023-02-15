package shopify

import "time"

// Fulfillment created when an order is fulfilled by a warehouse
type Fulfillment struct {
	// ID is the ID of the fulfillment
	ID int64
	// OrderID is the unique numeric identifier for the order.
	OrderID int64
	// The name of the tracking company.
	TrackingCompany string
	// TrackingNumber is a tracking number provided by the shipping company.
	TrackingNumber string
	// TrackingNumbers is a list of tracking numbers, provided by the shipping company.
	TrackingNumbers []string
	// The URL of the tracking page for the fulfillment.
	TrackingURL string
	// The URLs of tracking pages for the fulfillment.
	TrackingURLs []string
	// Status is the status of the fulfillment.
	/*
		Valid values:

		- pending: The fulfillment is pending.
		- open: The fulfillment has been acknowledged by the service and is in processing.
		- success: The fulfillment was successful.
		- cancelled: The fulfillment was cancelled.
		- error: There was an error with the fulfillment request.
		- failure: The fulfillment request failed.
	*/
	Status string
	// CreatedAt is the date and time when the fulfillment was created.
	CreatedAt time.Time
	// UpdatedAt is the date and time when the fulfillment was last modified.
	UpdatedAt time.Time
	// NotifyCustomer is a write only property setting whether or not changes to the fulfillment should trigger notifications to the customer.
	NotifyCustomer bool
	// ShipmentStatus is the current shipment status of the fulfillment.
	/*
		Valid values:

		- label_printed: A label for the shipment was purchased and printed.
		- label_purchased: A label for the shipment was purchased, but not printed.
		- attempted_delivery: Delivery of the shipment was attempted, but unable to be completed.
		- ready_for_pickup: The shipment is ready for pickup at a shipping depot.
		- confirmed: The carrier is aware of the shipment, but hasn't received it yet.
		- in_transit: The shipment is being transported between shipping facilities on the way to its destination.
		- out_for_delivery: The shipment is being delivered to its final destination.
		- delivered: The shipment was successfully delivered.
		- failure: 	Something went wrong when pulling tracking information for the shipment, such as the tracking number was invalid or the shipment was canceled.
	*/
	ShipmentStatus string
	// LocationID is the unique identifier of the location that the fulfillment should be processed for.
	LocationID int64
	// LineItems is a collection of historical records of each item in the fulfillment.
	LineItems LineItems
	// FulfillmentOrders is a required write only property describing the line items and fulfillment orders that this fulfillment fulfills.
	//
	// At least one fulfillment order is required when creating a fulfillment. Unless a fulfillment order has their LineItems field
	// assigned with at least one line item and quantity, all line items of the specified fulfillment order will be fulfilled.
	LineItemsByFulfillmentOrder FulfillmentOrders
}

// FulfillmentRepository manages fulfillments
type FulfillmentRepository interface {
	// Create creates a new fulfillment against a number of fulfillment orders, with optional tracking information and
	// customer notification
	//
	// LineItemsByFulfillmentOrder must be populated with at least one fulfillment order
	Create(fulfillment Fulfillment) (Fulfillment, error)
	// UpdateTracking updates the tracking details of a fulfillment
	UpdateTracking(fulfillment Fulfillment) (Fulfillment, error)
	// Cancel cancels a fulfillment
	Cancel(fulfillmentID int64) error
}
