package shopify

type Tax struct {
	// Rate is the percentage of the cost after discounts that should be charged as tax e.g. 0.2 for a 20% tax rate
	Rate float64
	// Title is the name of this kind of tax
	Title string
}
