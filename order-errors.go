package shopify

import "fmt"

// ErrOrderNotFound is thrown when no order is found with the ID
type ErrOrderNotFound struct {
	id int64
}

func (err ErrOrderNotFound) Error() string {
	return fmt.Sprintf("order %v not found", err.id)
}

// NewErrOrderNotFound builds the error
func NewErrOrderNotFound(id int64) ErrOrderNotFound {
	return ErrOrderNotFound{
		id,
	}
}
