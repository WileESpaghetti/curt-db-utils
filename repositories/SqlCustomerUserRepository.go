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

func (repo SqlCustomerRepository) getByEmail(email string) (c models.CustomerUser, err error) {
	customerUserFromEmail := `select customer.id, customer.email
				from CustomerUser as customer
				where UPPER(customer.email) = UPPER(?)
				`

	stmt, err := repo.session.Prepare(customerUserFromEmail)
	if err != nil {
		return c, err
	}
	defer stmt.Close()

	customerUser := models.CustomerUser{"", ""}
	result := stmt.QueryRow(email)
	err = result.Scan(&customerUser.Id, &customerUser.Email)

	return customerUser, err
}
