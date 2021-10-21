package shopify

import "time"

// Customer is a shopify customer
type Customer struct {
	// ID is the ID of the customer.
	ID int64
	// Email is the unique email address of the customer.
	Email string
	// FirstName is the customer's first name.
	FirstName string
	// LastName is the customer's last name.
	LastName string
	// CreatedAt is the date and time when the customer was created.
	CreatedAt time.Time
	// UpdatedAt is the date and time (ISO 8601 format) when the customer information was last updated.
	UpdatedAt time.Time
}

// CustomerRepository manages customers
type CustomerRepository interface {
	// Orders retrieves a list of orders belonging to a customer
	Orders(id int64) (Orders, error)
}
