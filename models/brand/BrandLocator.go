package brand

// Placeholder until we can get "*Store" classes to inherit from a more generic Store type
type BrandLocator interface {
	//Add(brand Brand) (err error)
	Get(id int) (brand Brand, err error)
	//Update(brand Brand) (err error)
	//Remove(brand Brand) (err error)
	All() (brand []Brand, err error)
}
