package shopify

// Shop represents a shopify shop
/*
	All shopify development documentation can be found at
	https://shopify.dev
*/
type Shop interface {
	// Orders are the orders in the shop
	Orders() OrderRepository
	// Fulfillments - you can use the Fulfillment resource to view, create, modify, or delete an order's or fulfillment order's fulfillments.
	Fulfillments() FulfillmentRepository
}
