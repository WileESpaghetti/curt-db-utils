package repositories

import (
	"github.com/WileESpaghetti/curt-db-utils/models"
)

type MongoCustomerRepository struct {
}

func (repo MongoCustomerRepository) getByApiKey(apiKey) models.Customer {
	// FIXME not implemented
	return nil
}
