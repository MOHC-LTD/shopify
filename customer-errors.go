package shopify

import "fmt"

// ErrCustomerNotFound is thrown when no customer is found with a given id
type ErrCustomerNotFound struct {
	id int64
}

func (err ErrCustomerNotFound) Error() string {
	return fmt.Sprintf("customer %v not found", err.id)
}

// NewErrCustomerNotFound builds the error
func NewErrCustomerNotFound(id int64) ErrCustomerNotFound {
	return ErrCustomerNotFound{
		id,
	}
}

// ErrCustomerSearchNotFound is thrown when no customer is found for the given search query
type ErrCustomerSearchNotFound struct {
	query string
}

func (err ErrCustomerSearchNotFound) Error() string {
	return fmt.Sprintf("customer from query `%v` was not found", err.query)
}

// ErrCustomerSearchNotFound builds the error
func NewErrCustomerSearchNotFound(query string) ErrCustomerSearchNotFound {
	return ErrCustomerSearchNotFound{
		query,
	}
}
