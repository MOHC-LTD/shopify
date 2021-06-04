package shopify

const (
	// FulfillmentShipmentStatusLabelPrinted -  A label for the shipment was purchased and printed
	FulfillmentShipmentStatusLabelPrinted = "label_printed"
	// FulfillmentShipmentStatusLabelPurchased - A label for the shipment was purchased, but not printed
	FulfillmentShipmentStatusLabelPurchased = "label_purchased"
	// FulfillmentShipmentStatusAttemptedDelivery - Delivery of the shipment was attempted, but unable to be completed
	FulfillmentShipmentStatusAttemptedDelivery = "attempted_delivery"
	// FulfillmentShipmentStatusReadyForPickup - The shipment is ready for pickup at a shipping depot
	FulfillmentShipmentStatusReadyForPickup = "ready_for_pickup"
	// FulfillmentShipmentStatusConfirmed - The carrier is aware of the shipment, but hasn't received it yet
	FulfillmentShipmentStatusConfirmed = "confirmed"
	// FulfillmentShipmentStatusInTransit - The shipment is being transported between shipping facilities on the way to its destination
	FulfillmentShipmentStatusInTransit = "in_transit"
	// FulfillmentShipmentStatusOutForDelivery - The shipment is being delivered to its final destination
	FulfillmentShipmentStatusOutForDelivery = "out_for_delivery"
	// FulfillmentShipmentStatusDelivered - The shipment was succesfully delivered
	FulfillmentShipmentStatusDelivered = "delivered"
	// FulfillmentShipmentStatusFailure - Something went wrong when pulling tracking information for the shipment, such as the tracking number was invalid or the shipment was canceled
	FulfillmentShipmentStatusFailure = "failure"
)

const (
	// FulfillmentStatusPending - The fulfillment is pending
	FulfillmentStatusPending = "pending"
	// FulfillmentStatusOpen - The fulfillment has been acknowledged by the service and is in processing
	FulfillmentStatusOpen = "open"
	// FulfillmentStatusSuccess - The fulfillment was successful
	FulfillmentStatusSuccess = "success"
	// FulfillmentStatusCancelled - The fulfillment was cancelled
	FulfillmentStatusCancelled = "cancelled"
	// FulfillmentStatusError - There was an error with the fulfillment request
	FulfillmentStatusError = "error"
	// FulfillmentStatusFailure - The fulfillment request failed
	FulfillmentStatusFailure = "failure"
)
