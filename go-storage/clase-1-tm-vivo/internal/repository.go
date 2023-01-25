package products

import "errors"

var (
	ErrNotFound   = errors.New("product not found")
	ErrInternal   = errors.New("an internal error")
	ErrDuplicated = errors.New("duplicated product")
)

type Repository interface {
	Get(id int) (*Product, error)
	Store(product *Product) error
	Update(product *Product) error
	Delete(id int) error
}
