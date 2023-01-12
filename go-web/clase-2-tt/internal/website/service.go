package website

import (
	"errors"
	"rest/internal/domain"
)

var (
	ErrAlreadyExist = errors.New("already exist")
)

// controller
type Service interface {
	Get() ([]domain.Website, error)
	GetByID(id int) (domain.Website, error)
	Create(url, host, category string, protected bool) (domain.Website, error)
}

func NewService(rp Repository) Service {
	return &service{rp: rp}
}

type service struct {
	// repo
	rp Repository

	// external api's
	// ...
}

// read
func (sv *service) Get() ([]domain.Website, error) {
	return sv.rp.Get()
}
func (sv *service) GetByID(id int) (domain.Website, error) {
	return sv.rp.GetByID(id)
}

// write
func (sv *service) Create(url, host, category string, protected bool) (domain.Website, error) {
	if sv.rp.ExistURL(url) {
		return domain.Website{}, ErrAlreadyExist
	}

	ws := domain.Website{
		URL: url,
		Host: host,
		Category: category,
		Protected: protected,
	}
	lastID, err := sv.rp.Create(ws)
	if err != nil {
		return domain.Website{}, err
	}

	ws.ID = lastID

	return ws, nil
}