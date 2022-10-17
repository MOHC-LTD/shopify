package shopify

// Address is an address
type Address struct {
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
	// FirstName - the first name of the person associated with the address
	FirstName string
	// Latitude - the latitude of the address
	Latitude float64
	// Longitude - the longitude of the address
	Longitude float64
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
