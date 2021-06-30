package shopify

// TaxLines is a collection of tax lines
type TaxLines []TaxLine

// TaxLine is a tax line
type TaxLine struct {
	// Title is the name of this kind of tax line
	Title string
	// Rate is the percentage of the cost after discounts that should be charged as tax e.g. 0.2 for a 20% tax rate
	Rate float32
}
