package brand

type BrandStore struct {
	Repository BrandRepository
}

func (store BrandStore) Get(id int) (brand Brand, err error) {
	return store.Repository.GetById(id)
}

func (store BrandStore) All() (brand []Brand, err error) {
	return store.Repository.All()
}
