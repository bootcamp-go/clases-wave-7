package movies

import (
	"errors"
	"rest/internal/domain"
)

var (
	ErrNotFound = errors.New("item not found")
)

// controller
type Repository interface {
	// read
	Get() ([]domain.Movie, error)
	GetByID(id int) (domain.Movie, error)
}

func NewRepository(db *[]domain.Movie) Repository {
	return &repository{db: db}
}

type repository struct {
	db	*[]domain.Movie
}
// read
func (rp *repository) Get() ([]domain.Movie, error) {
	return *rp.db, nil
}
func (rp *repository) GetByID(id int) (domain.Movie, error) {
	for _, m := range *rp.db {
		if m.ID == id {
			return m, nil
		}
	}

	return domain.Movie{}, ErrNotFound
}