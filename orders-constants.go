package shopify

const (
	// OrderFulfillmentStatusFulfilled is the fulfillment status of an order that is fulfilled
	OrderFulfillmentStatusFulfilled = "fulfilled"
	// OrderFulfillmentStatusPartial is the fulfillment status of an order that is partially fulfilled
	OrderFulfillmentStatusPartial = "partial"
	// OrderFulfillmentStatusRestocked is the fulfillment status of an order that has been restocked
	OrderFulfillmentStatusRestocked = "restocked"
)

const (
	// OrderFinancialStatusPending is the financial status of an order that has a pending payment
	OrderFinancialStatusPending = "pending"
	// OrderFinancialStatusAuthorized is the financial status of an order that has an authorized payment
	OrderFinancialStatusAuthorized = "auhthorized"
	// OrderFinancialStatusPartiallyPaid is the financial status of an order that has partially been payed for
	OrderFinancialStatusPartiallyPaid = "partially_paid"
	// OrderFinancialStatusPaid is the financial status of an order that has been payed for
	OrderFinancialStatusPaid = "paid"
	// OrderFinancialStatusPartiallyRefunded is the financial status of an order that has been partially refunded
	OrderFinancialStatusPartiallyRefunded = "partially_refunded"
	// OrderFinancialStatusRefunded is the financial status of an order that has been refunded
	OrderFinancialStatusRefunded = "refunded"
	// OrderFinancialStatusVoided is the financial status of an order that has had its payment voided
	OrderFinancialStatusVoided = "voided"
)

const (
	// OrderQueryFulfillmentStatusShipped shows orders that have been shipped
	OrderQueryFulfillmentStatusShipped = "shipped"
	// OrderQueryFulfillmentStatusPartial shows partially shipped orders
	OrderQueryFulfillmentStatusPartial = "partial"
	// OrderQueryFulfillmentStatusUnshipped shows orders that have not yet been shipped
	OrderQueryFulfillmentStatusUnshipped = "unshipped"
	// OrderQueryFulfillmentStatusAny shows orders of any fulfillment status
	OrderQueryFulfillmentStatusAny = "any"
	// OrderQueryFulfillmentStatusUnfulfilled shows orders that have not been fulfilled or have been partially fulfilled
	OrderQueryFulfillmentStatusUnfulfilled = "unfulfilled"
)

const (
	// OrderQueryFinancialStatusPending shows only pending orders
	OrderQueryFinancialStatusPending = "pending"
	// OrderQueryFinancialStatusAuthorized shows only authorized orders
	OrderQueryFinancialStatusAuthorized = "auhthorized"
	// OrderQueryFinancialStatusPartiallyPaid shows only partially paid orders
	OrderQueryFinancialStatusPartiallyPaid = "partially_paid"
	// OrderQueryFinancialStatusPaid shows only paid orders
	OrderQueryFinancialStatusPaid = "paid"
	// OrderQueryFinancialStatusPartiallyRefunded shows only partially refunded orders
	OrderQueryFinancialStatusPartiallyRefunded = "partially_refunded"
	// OrderQueryFinancialStatusRefunded shows only refunded orders
	OrderQueryFinancialStatusRefunded = "refunded"
	// OrderQueryFinancialStatusVoided shows only voided orders
	OrderQueryFinancialStatusVoided = "voided"
	// OrderQueryFinancialStatusAny shows orders of any financial status
	OrderQueryFinancialStatusAny = "any"
	// OrderQueryFinancialStatusUnpaid shows authorized and partially paid orders
	OrderQueryFinancialStatusUnpaid = "unpaid"
)

const (
	// OrderQueryStatusOpen show only open orders
	OrderQueryStatusOpen = "open"
	// OrderQueryStatusClosed show only closed orders
	OrderQueryStatusClosed = "closed"
	// OrderQueryStatusCancelled show only canceled orders
	OrderQueryStatusCancelled = "cancelled"
	// OrderQueryStatusAny show orders of any status, including archived orders
	OrderQueryStatusAny = "any"
)
