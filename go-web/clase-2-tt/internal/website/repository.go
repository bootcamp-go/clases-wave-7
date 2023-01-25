package website

import (
	"errors"
	"fmt"
	"rest/internal/domain"
)

var (
	ErrNotFound = errors.New("item not found")
)

type Repository interface {
	// read
	Get() ([]domain.Website, error)
	GetByID(id int) (domain.Website, error)
	ExistURL(url string) bool
	// write
	Create(domain.Website) (int, error)
}

type repository struct {
	db *[]domain.Website
	// config
	lastID	int
}

func NewRepository(db *[]domain.Website, lastID	int) Repository {
	return &repository{db: db, lastID: lastID}
}

// read
func (r *repository) Get() ([]domain.Website, error) {
	return *r.db, nil
}
func (r *repository) GetByID(id int) (domain.Website, error) {
	for _, w := range *r.db {
		if w.ID == id {
			return w, nil
		}
	}

	return domain.Website{}, fmt.Errorf("%w. %s", ErrNotFound, "website does not exist")
}
func (r *repository) ExistURL(url string) bool {
	for _, w := range *r.db {
		if w.URL == url {
			return true
		}
	}

	return false
}

// write
func (r *repository) Create(website domain.Website) (int, error) {
	r.lastID++
	website.ID = r.lastID
	*r.db = append(*r.db, website)

	return r.lastID, nil
}