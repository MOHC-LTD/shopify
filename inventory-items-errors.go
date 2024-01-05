package shopify

import "fmt"

// ErrInventoryItemNotFound is thrown when no inventory item is found with a given id
type ErrInventoryItemNotFound struct {
	id int64
}

func (err ErrInventoryItemNotFound) Error() string {
	return fmt.Sprintf("inventory item %v not found", err.id)
}

// NewErrInventoryItemNotFound builds the error
func NewErrInventoryItemNotFound(id int64) ErrInventoryItemNotFound {
	return ErrInventoryItemNotFound{
		id,
	}
}
