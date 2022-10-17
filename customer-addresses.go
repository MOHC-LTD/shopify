package shopify

// CustomerAddressRepository maintains the customer addresses of a shop.
type CustomerAddressRepository interface {
	// Create creates a new address against the specified customer
	Create(id int64, address Address) (Address, error)
	// Delete deletes the specified address for the specified customer
	Delete(id int64, addressID int64) error
}
