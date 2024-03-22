package shopify

import "fmt"

// ErrTransactionNotFound is thrown when no transaction is found with the ID
type ErrTransactionNotFound struct {
	id int64
}

func (err ErrTransactionNotFound) Error() string {
	return fmt.Sprintf("transaction %v not found", err.id)
}

// NewErrTransactionNotFound builds the error
func NewErrTransactionNotFound(id int64) ErrTransactionNotFound {
	return ErrTransactionNotFound{
		id,
	}
}
