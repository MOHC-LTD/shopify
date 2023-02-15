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
	// FulfillmentOrders represent groups of line items that can be fulfilled from the same location.
	FulfillmentOrders() FulfillmentOrderRepository
	// Variants represents options or versions that make up an individual product.
	Variants() VariantRepository
	// Products are the items that are sold by the shop.
	Products() ProductRepository
	// InventoryLevels are the amount of stock available for each product
	InventoryLevels() InventoryLevelRepository
	// Collections are groups of products that are sold by the shop
	Collections() CollectionRepository
	// ProductImages are the images associated with products
	ProductImages() ProductImageRepository
	// Metafields are the metafields associated with the shop
	Metafields() MetafieldRepository
}
