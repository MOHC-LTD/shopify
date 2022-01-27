package shopify

// DiscountApplications is a collection of discount applications
type DiscountApplications []DiscountApplication

// DiscountApplication is an application of a discount on an order
type DiscountApplication struct {
	// AllocationMethod is the method by which the discount application has been allocated to entitled lines
	/*
		Valid values:

		- across: The value is spread across all entitled lines
		- each: The value is applied onto every entitled line
		- one: The value is applied onto a single line
	*/
	AllocationMethod string
	// Code is the code that was used for a discount code application
	Code string
	// Type is the type of the discount application
	/*
		Valid values:

		- automatic: The discount was applied automatically
		- discount_code: The discount was applied by a discount code
		- manual: The discount was manually applied by the merchant
		- script: The discount was applied by a Shopify Script
	*/
	Type string
	// Title is the title of the discount application as defined by the merchant
	Title string
	// Value is the value of the discount application as a decimal
	Value string
	// ValueType is the type of the value
	/*
		Valid values:

		- fixed_amount: A fixed amount discount value in the currency of the order.
		- percentage: A percentage discount value.
	*/
	ValueType string
	// Description is the description of the discount application as defined by the merchant or the Shopify Script
	Description string
	// TargetType is the type of line on the order that is the discount is applicable on
	/*
		Valid values:

		- line_item: The discount applies to line items
		- shipping_line: The discount applies to shipping lines
	*/
	TargetType string
	// TargetSelection is the lines of the order, as specified by TargetType, that the discount is allocated over
	/*
		Valid values:

		- all: The discount is allocated onto all lines
		- entitled: The discount is allocated only onto lines it is entitled for
		- explicit: The discount is allocated onto explicitly selected lines
	*/
	TargetSelection string
}

// DiscountAllocations is a collection of discount allocations
type DiscountAllocations []DiscountAllocation

// DiscountAllocation is an ordered list of amounts allocated by discount applications
type DiscountAllocation struct {
	// Amount is the discount amount allocated to the line in the shop currency
	Amount string
	// AmountSet is the discount amount allocated to the line in shop and presentment currencies
	AmountSet PriceSet
	// DiscountApplicationIndex is the index of the associated DiscountApplication in the order's DiscountApplications list
	DiscountApplicationIndex int
}
