package brand

// TODO validate fields and throw error if incorrect. Not sure what level this should be
// Low level data access layer for Brands
type BrandRepository interface {
	GetById(id int) (brand Brand, err error)
	SaveNew(brand *Brand) (err error)
	All() ([]Brand, error)
}
