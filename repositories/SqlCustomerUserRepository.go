package repositories

import (
	"github.com/WileESpaghetti/curt-db-utils/models"
	"database/sql"
)

type SqlCustomerRepository struct {
	Session *sql.DB
}

func (repo SqlCustomerRepository) GetByApiKey(apiKey string) (c models.CustomerUser, err error) {
	// FIXME not implemented
	return models.CustomerUser{}, nil
}

func (repo SqlCustomerRepository) GetByEmail(email string) (c models.CustomerUser, err error) {
	customerUserFromEmail := `select customer.id, customer.email
				from CustomerUser as customer
				where UPPER(customer.email) = UPPER(?)
				`

	stmt, err := repo.Session.Prepare(customerUserFromEmail)
	if err != nil {
		return c, err
	}
	defer stmt.Close()

	customerUser := models.CustomerUser{}
	result := stmt.QueryRow(email)
	err = result.Scan(&customerUser.Id, &customerUser.Email)

	return customerUser, err
}
