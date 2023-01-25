package movies

import "rest/internal/domain"

// controller
type Service interface {
	// read
	Get() ([]domain.Movie, error)
	GetByID(id int) (domain.Movie, error)
}

func NewService(rp Repository) Service {
	return &service{rp: rp}
}

type service struct {
	rp Repository
}
// read
func (sv *service) Get() ([]domain.Movie, error) {
	return sv.rp.Get()
}
func (sv *service) GetByID(id int) (domain.Movie, error) {
	return sv.rp.GetByID(id)
}