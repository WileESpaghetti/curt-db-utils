package models

type BrandRepository interface {
	GetById(id int) (brand Brand, err error)
}
