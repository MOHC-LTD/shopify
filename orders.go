package shopify

import (
	"time"
)

// Orders is a collection of orders
type Orders []Order

// Order is a customer's completed request to purchase one or more products from a shop.
type Order struct {
	// BillingAddress is the mailing address associated with the payment method.
	BillingAddress Address
	// ClosedAt is the date and time when the order was closed.
	ClosedAt time.Time
	// CreatedAt is the autogenerated date and time when the order was created in Shopify. The value for this property cannot be changed.
	CreatedAt time.Time
	// Currency is the three-letter code (ISO 4217 format) for the shop currency.
	Currency string
	// CurrentTotalDiscounts is the current total discounts on the order in the shop currency. The value of this field reflects order edits, returns, and refunds.
	CurrentTotalDiscounts string
	// CurrentTotalDiscountsSet is the current total discounts on the order in shop and presentment currencies. The amount values associated with this field reflect order edits, returns, and refunds.
	CurrentTotalDiscountsSet PriceSet
	// CurrentTotalPrice is the sum of all line item prices, discounts, shipping, taxes, and tips in the shop currency. It is always positive.
	CurrentTotalPrice string
	// CurrentTotalPriceSet is the current total price of the order in shop and presentment currencies. The amount values associated with this field reflect order edits, returns, and refunds.
	CurrentTotalPriceSet PriceSet
	// CurrentSubtotalPrice is the current subtotal price of the order in the shop currency. The value of this field reflects order edits, returns, and refunds.
	CurrentSubtotalPrice string
	// CurrentSubtotalPriceSet is the current subtotal price of the order in shop and presentment currencies. The amount values associated with this field reflect order edits, returns, and refunds.
	CurrentSubtotalPriceSet PriceSet
	// CurrentTotalTax is the current total taxes charged on the order in the shop currency. The value of this field reflects order edits, returns, or refunds.
	CurrentTotalTax string
	// CurrentTotalTaxSet is the current total taxes charged on the order in shop and presentment currencies. The amount values associated with this field reflect order edits, returns, and refunds.
	CurrentTotalTaxSet PriceSet
	// Customer is information about the customer.
	Customer Customer
	// DiscountApplications is an ordered list of discount applications
	DiscountApplications DiscountApplications
	// Email is the customer's email address.
	Email string
	// FinancialStatus is the status of payments associated with the order. Can only be set when the order is created.
	/*
		Valid values:

		- pending: The payments are pending. Payment might fail in this state. Check again to confirm whether the payments have been paid successfully.
		- authorized: The payments have been authorized.
		- partially_paid: The order have been partially paid.
		- paid: The payments have been paid.
		- partially_refunded: The payments have been partially refunded.
		- refunded: The payments have been refunded.
		- voided: The payments have been voided.
	*/
	FinancialStatus string
	// Fulfillments is a list of fulfillments associated with the order.
	Fulfillments []Fulfillment
	// FulfillmentStatus - The order's status in terms of fulfilled line items.
	/*
		Valid values:

		- fulfilled: Every line item in the order has been fulfilled.
		- "": None of the line items in the order have been fulfilled.
		- partial: At least one line item in the order has been fulfilled.
		- restocked: Every line item in the order has been restocked and the order canceled.
	*/
	FulfillmentStatus string
	// ID is the ID of the order
	ID int64
	// LineItems is a list of line item objects, each containing information about an item in the order.
	LineItems LineItems
	// The order name, generated by combining the number property with the order prefix and suffix.
	Name string
	// Extra information that is added to the order. Appears in the Additional details section of an order details
	// page. Each array entry must contain a hash with name and value keys.
	NoteAttributes NoteAttributes
	// OrderNumber tells you that this is the shop's nth order.
	OrderNumber int
	// PresentmentCurrency is the presentment currency that was used to display prices to the customer.
	PresentmentCurrency string
	// ProcessedAt is the date and time the order was created at
	ProcessedAt time.Time
	// ShippingAddress is the mailing address to where the order will be shipped.
	ShippingAddress Address
	// ShippingLines is a collection of details about the shipping method used.
	ShippingLines ShippingLines
	// SubtotalPrice is the price of the order in the shop currency after discounts but before shipping, duties, taxes, and tips.
	SubtotalPrice string
	// SubtotalPriceSet is the subtotal of the order in shop and presentment currencies after discounts but before shipping, duties, taxes, and tips.
	SubtotalPriceSet PriceSet
	// Tags are tags attached to the order, formatted as a string of comma-separated values. Tags are additional short descriptors, commonly used for filtering and searching. Each individual tag is limited to 40 characters in length.
	Tags Tags
	// TotalDiscounts is the total discounts applied to the price of the order in the shop currency.
	TotalDiscounts string
	// TotalDiscountsSet is the total discounts applied to the price of the order in shop and presentment currencies.
	TotalDiscountsSet PriceSet
	// TotalLineIemsPrice is the sum of all line item prices in the shop currency.
	TotalLineItemsPrice string
	// TotalLineItemsPriceSet is the total of all line item prices in shop and presentment currencies.
	TotalLineItemsPriceSet PriceSet
	// TotalPrice is the sum of all line item prices, discounts, shipping, taxes, and tips in the shop currency. It is always positive.
	TotalPrice string
	// TotalPriceSet is the total price of the order in shop and presentment currencies.
	TotalPriceSet PriceSet
	// TotalTax is the sum of all the taxes applied to the order in the shop currency. It is always positive.
	TotalTax string
	// TotalTaxSet is the total tax applied to the order in shop and presentment currencies.
	TotalTaxSet PriceSet
	// SourceName is the source of the checkout
	SourceName string
	// UpdatedAt is the date and time when the order was last modified.
	UpdatedAt time.Time
}

// NoteAttributes represents a collection of note attributes.
type NoteAttributes []NoteAttribute

// NoteAttribute represents a single note attribute.
type NoteAttribute struct {
	// Name represents the name of the attribute.
	Name string
	// Value represents the value of the attribute.
	Value string
}

// OrderQuery are properties that can be used to filter the returned orders
// See https://shopify.dev/docs/admin-api/rest/reference/orders/order#index-2021-04
type OrderQuery struct {
	/*
		Show orders created at or after date.
	*/
	CreatedAtMin time.Time
	/*
		Filter orders by their fulfillment status.

		(default: any)
		- shipped: Show orders that have been shipped. Returns orders with fulfillment_status of fulfilled.
		- partial: Show partially shipped orders.
		- unshipped: Show orders that have not yet been shipped. Returns orders with fulfillment_status of null.
		- any: Show orders of any fulfillment status.
		- unfulfilled: Returns orders with fulfillment_status of null or partial.
	*/
	FulfillmentStatus string
	/*
		Filter orders by their financial status.

		(default: any)
		- authorized: Show only authorized orders
		- pending: Show only pending orders
		- paid: Show only paid orders
		- partially_paid: Show only partially paid orders
		- refunded: Show only refunded orders
		- voided: Show only voided orders
		- partially_refunded: Show only partially refunded orders
		- any: Show orders of any financial status.
		- unpaid: Show authorized and partially paid orders.
	*/
	FinancialStatus string
	/*
		The maximum number of results to show on a page (250 is the current max).
	*/
	Limit int
	/*
		Filter orders by their status.

		(default: open)
		- open: Show only open orders.
		- closed: Show only closed orders.
		- cancelled: Show only canceled orders.
		- any: Show orders of any status, including archived orders.
	*/
	Status string
	/*
		Show orders after the specified ID.
	*/
	SinceID int64
	/*
		Retrieve only orders specified by a list of order IDs.
	*/
	IDs []int64
}

// OrderRepository maintains the orders in the shop
type OrderRepository interface {
	// List gets all the orders
	List(query OrderQuery) (Orders, error)
	// Get gets an order
	Get(id int64) (Order, error)
	// Close closes an order
	Close(id int64) error
	// Open opens an order
	Open(id int64) (Order, error)
	// Create creates a new order
	Create(order Order) (Order, error)
	// Update updates a single order
	Update(order Order) (Order, error)
	// CreateMetafield creates a single metafield for an order
	CreateMetafield(orderID int64, metafield Metafield) (Metafield, error)
	// UpdateMetafield updates a single metafield
	UpdateMetafield(orderID int64, metafield Metafield) (Metafield, error)
}
