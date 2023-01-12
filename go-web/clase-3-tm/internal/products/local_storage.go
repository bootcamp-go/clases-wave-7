package products

import "github.com/google/uuid"

// LocalStorage implements the Storage interface using a local slice.
type LocalStorage struct {
	Products []Product
}

// Store a product in the storage.
func (storage *LocalStorage) Store(product *Product) (err error) {
	// Create an identifier for the product.
	product.ID = uuid.New().String()

	// Add the product to the slice.
	storage.Products = append(storage.Products, *product)
	return
}

// GetAll returns all the products in the storage.
func (storage *LocalStorage) GetAll() (result []Product, err error) {
	result = storage.Products
	return
}

// Update a product in the storage.
func (storage *LocalStorage) Update(id string, product *Product) (err error) {
	var updated bool

	// Find the product in the slice.
	for index := range storage.Products {
		if storage.Products[index].ID != id {
			continue
		}

		// Update the product.
		storage.Products[index] = *product
		updated = true
		break
	}

	// Check if the product was updated.
	if !updated {
		err = ErrProductNotFound
		return
	}

	return
}

// GetByID returns a product by its ID.
func (storage *LocalStorage) GetByID(id string) (result *Product, err error) {
	// Find the product in the slice.
	for index := range storage.Products {
		if storage.Products[index].ID != id {
			continue
		}

		// Return the product.
		result = &storage.Products[index]
		break
	}

	// Check if the product was found.
	if result == nil {
		err = ErrProductNotFound
		return
	}

	return
}

// Delete a product from the storage.
func (storage *LocalStorage) Delete(id string) (err error) {
	var found bool

	// Find the product in the slice.
	for index := range storage.Products {
		if storage.Products[index].ID != id {
			continue
		}

		// Delete the product.
		storage.Products = append(storage.Products[:index], storage.Products[index+1:]...)
		found = true
		break
	}

	// Check if the product was deleted.
	if !found {
		err = ErrProductNotFound
		return
	}

	return
}
