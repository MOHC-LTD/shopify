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
	// FulfillmentEvents represent tracking events that belong to a fulfillment of one or more items in an order.
	FulfillmentEvents() FulfillmentEventRepository
	// Variants represents options or versions of that make up an individual product.
	Variants() VariantRepository
}
