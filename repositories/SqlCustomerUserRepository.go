package repositories

import (
	"github.com/WileESpaghetti/curt-db-utils/models"
	"database/sql"
)

type SqlCustomerRepository struct {
	session *sql.DB
}

func (repo SqlCustomerRepository) getByApiKey(apiKey string) models.CustomerUser {
	// FIXME not implemented
	return nil
}

func (repo SqlCustomerRepository) getByEmail(email string) models.CustomerUser {
	// FIXME not implemented
	return nil
}
