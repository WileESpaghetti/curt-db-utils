package brand

// TODO Eventually it would be nice if this could implement from a Store interface
type BrandStore struct {
	Repository BrandRepository
}

func (store BrandStore) Get(id int) (brand Brand, err error) {
	return store.Repository.GetById(id)
}

func (store BrandStore) All() (brand []Brand, err error) {
	return store.Repository.All()
}
