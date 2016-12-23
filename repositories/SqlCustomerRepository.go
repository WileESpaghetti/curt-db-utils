package repositories

import (
	"github.com/WileESpaghetti/curt-db-utils/models"
)

type SqlCustomerRepository struct {
}

func (repo SqlCustomerRepository) getByApiKey(apiKey) models.Customer {
	// FIXME not implemented
	return nil
}
