package shopify

import (
	"errors"
	"fmt"
)

// ErrVariantNotFound is thrown when no variant is found with the id
type ErrVariantNotFound struct {
	id int64
}

func (err ErrVariantNotFound) Error() string {
	return fmt.Sprintf("variant %v not found", err.id)
}

// NewErrVariantNotFound builds the error
func NewErrVariantNotFound(id int64) ErrVariantNotFound {
	return ErrVariantNotFound{
		id,
	}
}

// ErrVariantNotFoundByPosition is thrown when no variant is found with the id
var ErrVariantNotFoundByPosition = errors.New("variant not found")
