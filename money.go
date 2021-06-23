package shopify

type Money struct {
	// Amount is the amount of money without the currency unit e.g. "5.00"
	Amount string
	// CurrencyCode is the 3-letter code that represents a currency e.g EUR or GBP
	CurrencyCode string
}
