package services

import (
	"errors"
	"fmt"
	"rest/services/models"
)

var (
	ErrAlreadyExist = errors.New("error: item already exist")
)

var websites = []models.WebSite{
	{ID: 1, URL: "https://www.google.com", Host: "google", Category: "search", Protected: false},
	{ID: 2, URL: "https://www.bing.com", Host: "bing", Category: "search", Protected: true},
	{ID: 3, URL: "https://www.mercadolibre.com", Host: "meli", Category: "e-commerce", Protected: false},
	{ID: 4, URL: "https://www.mercadopago.com", Host: "meli", Category: "finance", Protected: false},
}
var lastID = 4

// read
func Get() []models.WebSite {
	return websites
}
func GetByID() {

}
func ExistURL(url string) bool {
	for _, w := range websites {
		if w.URL == url {
			return true
		}
	}

	return false
}

// write
func Create(url, host, category string, protected bool) (models.WebSite, error) {
	// validations
	if ExistURL(url) {
		return models.WebSite{}, fmt.Errorf("%w. %s", ErrAlreadyExist, "url not unique")
	}
	
	lastID++
	website := models.WebSite{
		ID: lastID,
		URL: url,
		Host: host,
		Category: category,
		Protected: protected,
	}
	
	websites = append(websites, website)
	return website, nil
}