package shopify

// CustomerAddressRepository maintains the customer addresses of a shop.
type CustomerAddressRepository interface {
	// Delete deletes the specified address for the specified customer
	Delete(id int64, addressID int64) error
}
