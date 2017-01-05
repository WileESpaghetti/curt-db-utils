package repositories

import (
	"github.com/WileESpaghetti/curt-db-utils/models"
	"database/sql"
	"fmt"
)

type SqlCustomerRepository struct {
	Session *sql.DB
}

func (repo SqlCustomerRepository) GetByApiKey(apiKey string, apiKeyType string) (c models.CustomerUser, err error) {
	customerUserFromKey := `select
		cu.id,
		cu.email
	from CustomerUser as cu
	join ApiKey as ak
		on cu.id = ak.user_id
	join ApiKeyType as akt
		on ak.type_id = akt.id
	where
		UPPER(akt.type) = ? &&
		UPPER(ak.api_key) = UPPER(?)
	limit 1
	`
	stmt, err := repo.Session.Prepare(customerUserFromKey)
	if err != nil {
		return c, err
	}
	defer stmt.Close()

	customerUser := models.CustomerUser{}
	result := stmt.QueryRow(apiKeyType, apiKey)
	err = result.Scan(&customerUser.Id, &customerUser.Email)

	if err != nil {
		err = fmt.Errorf("error: %s", "API Key or User does not exist")
	}

	return customerUser, err
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
