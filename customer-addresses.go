package shopify

// CustomerAddressRepository maintains the customer addresses of a shop.
type CustomerAddressRepository interface {
	// Create creates a new address against the specified customer
	Create(id int64, address CustomerAddress) (CustomerAddress, error)
	// Delete deletes the specified address for the specified customer
	Delete(id int64, addressID int64) error
}

type CustomerAddress struct {
	// Address1 - the street address
	Address1 string
	// Address2 - an optional additional field for the street address
	Address2 string
	// City - the city, town, or village
	City string
	// Company - the company of the person associated with this address
	Company string
	// Country - the name of the country
	Country string
	// CountryCode - the two-letter code (ISO 3166-1 format) for the country of the address
	CountryCode string
	// CountryName - the name of the country
	CountryName string
	// CustomerID - the ID of the customer which address is stored against.
	CustomerID uint64
	// Default - Whether this is the customer default address.
	Default bool
	// FirstName - the first name of the person associated with the address
	FirstName string
	// ID - the ID of the address
	ID uint64
	// LastName - the last name of the person associated with the address
	LastName string
	// Name - the full name of the person associated with the address
	Name string
	// Phone - the phone number at the address
	Phone string
	// Province - the name of the region (for example, province, state, or prefecture) of the address
	Province string
	// ProvinceCode - the two-letter abbreviation of the region of the address
	ProvinceCode string
	// Zip - the postal code (for example, zip, postcode or Eircode) of the address
	Zip string
}
