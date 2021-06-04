package shopify

const (
	// FulfillmentEventStatusLabelPrinted -  A label for the shipment was purchased and printed
	FulfillmentEventStatusLabelPrinted = "label_printed"
	// FulfillmentEventStatusLabelPurchased - A label for the shipment was purchased, but not printed
	FulfillmentEventStatusLabelPurchased = "label_purchased"
	// FulfillmentEventStatusAttemptedDelivery - Delivery of the shipment was attempted, but unable to be completed
	FulfillmentEventStatusAttemptedDelivery = "attempted_delivery"
	// FulfillmentEventStatusReadyForPickup - The shipment is ready for pickup at a shipping depot
	FulfillmentEventStatusReadyForPickup = "ready_for_pickup"
	// FulfillmentEventStatusPickedUp - The fulfillment was successfully picked up
	FulfillmentEventStatusPickedUp = "picked_ip"
	// FulfillmentEventStatusConfirmed - The carrier is aware of the shipment, but hasn't received it yet
	FulfillmentEventStatusConfirmed = "confirmed"
	// FulfillmentEventStatusInTransit - The shipment is being transported between shipping facilities on the way to its destination
	FulfillmentEventStatusInTransit = "in_transit"
	// FulfillmentEventStatusOutForDelivery - The shipment is being delivered to its final destination
	FulfillmentEventStatusOutForDelivery = "out_for_delivery"
	// FulfillmentEventStatusDelivered - The shipment was succesfully delivered
	FulfillmentEventStatusDelivered = "delivered"
	// FulfillmentEventStatusFailure - Something went wrong when pulling tracking information for the shipment, such as the tracking number was invalid or the shipment was canceled
	FulfillmentEventStatusFailure = "failure"
)
