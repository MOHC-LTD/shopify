package shopify

import (
	"net/url"
	"time"
)

// Customer is a shopify customer
type Customer struct {
	// ID is the ID of the customer.
	ID int64
	// Email is the unique email address of the customer.
	Email string
	// Phone is the unique phone number (E.164 format) for this customer.
	Phone string
	// FirstName is the customer's first name.
	FirstName string
	// LastName is the customer's last name.
	LastName string
	// Tags are the tags linked to the customer.
	Tags Tags
	// CreatedAt is the date and time when the customer was created.
	CreatedAt time.Time
	// UpdatedAt is the date and time (ISO 8601 format) when the customer information was last updated.
	UpdatedAt time.Time
}

// Customers is a collection of customers
type Customers []Customer

// CustomerSearchQuery contains all the supported search queries
type CustomerSearchQuery struct {
	Tag string
}

func (c *CustomerSearchQuery) String() string {
	params := url.Values{}

	if c.Tag != "" {
		params.Add("customer_tag:", c.Tag)
	}

	return params.Encode()
}

// CustomerRepository maintains the customers of a shop.
type CustomerRepository interface {
	// Update updates a single customer
	Update(customer Customer) (Customer, error)
	// Get gets customer with the provided id
	Get(id int64) (Customer, error)
	// GetByQuery gets the customers matching the query and returns the fields requested
	GetByQuery(fields []string, query CustomerSearchQuery) (Customers, error)
}
