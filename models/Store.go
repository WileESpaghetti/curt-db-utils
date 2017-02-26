package models

type Store interface {
	Get(id int) (interface{}, error)
	All() ([]interface{}, error)
}
