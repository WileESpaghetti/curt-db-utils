package brand

// Placeholder until we can get "*Store" classes to inherit from a more generic Store type
type BrandLocator interface {
	Get(id int) (brand Brand, err error)
	All() (brand []Brand, err error)
}
