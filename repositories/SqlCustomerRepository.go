package repositories

import (
	"github.com/WileESpaghetti/curt-db-utils/models"
)

type SqlCustomerRepository struct {
}

func (repo SqlCustomerRepository) getByApiKey(apiKey string) models.Customer {
	// FIXME not implemented
	return nil
}

func (repo SqlCustomerRepository) getByEmail(email string) models.Customer {
	// FIXME not implemented
	return nil
}
