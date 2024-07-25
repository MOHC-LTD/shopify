package shopify

import "fmt"

// LineItems is a collection of line items
type LineItems []LineItem

// LineItem is an item on an order
type LineItem struct {
	// ID is the ID of the line item
	ID int64
	// Price is the cost of the item before any discounts or tax are applied, in the store's currency
	Price string
	// PriceSet is the price of the line item in shop and presentment currencies.
	PriceSet PriceSet
	// ProductID is the ID of the product that the line item belongs to. Can be 0 if the original product associated with the order is deleted at a later date.
	ProductID int64
	// Quantity is the number of items that were purchased.
	Quantity int
	// SKU is the item's SKU (stock keeping unit).
	SKU string
	// Title is the title of the product.
	Title string
	// VariantID is the ID of the product variant.
	VariantID int64
	// VariantTitle is the title of the product variant.
	VariantTitle string
	// Name is the name of the product variant.
	Name string
	// TaxLines is a list of tax line objects, each of which details a tax applied to the item.
	TaxLines TaxLines
	// TotalDiscount is the total amount of the discount allocated to the line item in the shop currency.
	TotalDiscount string
	// TotalDiscountSet is the total amount allocated to the line item in the presentment currency.
	TotalDiscountSet PriceSet
	// DiscountAllocations is a list of the amounts allocated to the line item by discount applications.
	DiscountAllocations DiscountAllocations
	// Properties is an array of custom information for the item that has been added to the cart. Often used to provide product customization options.
	Properties []Property
}

// Property holds custom information for the item that has been added to the cart. Often used to provide product customization options.
type Property struct {
	// Name is the name of the property
	Name string
	// Value is the value of the property
	Value interface{}
}

// GetPropertyByName gets a line item property by its name
func (lineItem LineItem) GetPropertyByName(name string) Property {
	for _, property := range lineItem.Properties {
		if property.Name == name || property.Name == fmt.Sprintf("_%s", name) {
			return property
		}
	}
	return Property{}
}
