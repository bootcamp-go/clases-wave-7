package products

import "errors"

var (
	// ErrProductNotFound is the error returned when a product is not found.
	ErrProductNotFound = errors.New("product not found")
)

// Storage defines the behavior of a products storage.
type Storage interface {
	// Store a product in the storage.
	Store(product *Product) error

	// GetAll returns all the products in the storage.
	GetAll() ([]Product, error)

	// Update a product in the storage.
	Update(id string, product *Product) error

	// GetByID returns a product by its ID.
	GetByID(id string) (*Product, error)

	// Delete a product from the storage.
	Delete(id string) error
}
