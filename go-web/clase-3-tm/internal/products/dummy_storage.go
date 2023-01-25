package products

type DummyStorage struct {
	ProductOnStoreMethod *Product
	ErrOnStoreMethod     error

	ResultOnGetAllMethod []Product
	ErrOnGetAllMethod    error

	ProductOnUpdateMethod *Product
	ErrOnUpdateMethod     error

	ProductOnGetByIDMethod *Product
	ErrOnGetByIDMethod     error

	ErrOnDeleteMethod error
}

// Store a product in the storage.
func (storage *DummyStorage) Store(product *Product) error {
	product.ID = storage.ProductOnStoreMethod.ID
	return storage.ErrOnStoreMethod
}

// GetAll returns all the products in the storage.
func (storage *DummyStorage) GetAll() ([]Product, error) {
	return storage.ResultOnGetAllMethod, storage.ErrOnGetAllMethod
}

// Update a product in the storage.
func (storage *DummyStorage) Update(id string, product *Product) error {
	product = storage.ProductOnUpdateMethod
	return storage.ErrOnUpdateMethod
}

// GetByID returns a product by its ID.
func (storage *DummyStorage) GetByID(id string) (*Product, error) {
	return storage.ProductOnGetByIDMethod, storage.ErrOnGetByIDMethod
}

// Delete a product from the storage.
func (storage *DummyStorage) Delete(id string) error {
	return storage.ErrOnDeleteMethod
}
