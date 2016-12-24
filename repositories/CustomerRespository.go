package repositories

import "github.com/WileESpaghetti/curt-db-utils/models"

type CustomerUserRepository interface {
	getByApiKey(apiKey string) (c models.CustomerUser, err error)
	getByEmail(email string) (c models.CustomerUser, err error)
}
