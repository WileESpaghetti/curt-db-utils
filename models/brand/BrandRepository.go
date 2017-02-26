package brand

type BrandRepository interface {
	GetById(id int) (brand Brand, err error)
	SaveNew(brand *Brand) (err error)
	All() ([]Brand, error)
}
