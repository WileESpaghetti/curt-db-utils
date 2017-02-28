package brand

// TODO Eventually it would be nice if this could implement from a Store interface
type BrandStore struct {
	// Data access interface
	Repository BrandRepository
}

// Retrieve a stored Brand
func (store BrandStore) Get(id int) (brand Brand, err error) {
	return store.Repository.GetById(id)
}

// Retrieve all available Brands
func (store BrandStore) All() (brand []Brand, err error) {
	return store.Repository.All()
}
