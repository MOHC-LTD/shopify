package shopify

// PriceSet represents a price in shop and presentment currencies
type PriceSet struct {
	// ShopMoney is the price in shop currency
	ShopMoney Money
	// PresentmentMoney is the price in presentment currency
	PresentmentMoney Money
}

// Money is an amount of money in some currency
type Money struct {
	// Amount is the amount of money without the currency unit e.g. "5.00"
	Amount string
	// CurrencyCode is the 3-letter code that represents a currency e.g EUR or GBP
	CurrencyCode string
}
