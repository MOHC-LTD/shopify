package shopify

import "fmt"

// ErrProductNotFound is thrown when no product is found with the id
type ErrProductNotFound struct {
	id int64
}

func (err ErrProductNotFound) Error() string {
	return fmt.Sprintf("product %v not found", err.id)
}

// NewErrProductNotFound builds the error
func NewErrProductNotFound(id int64) ErrProductNotFound {
	return ErrProductNotFound{
		id,
	}
}
