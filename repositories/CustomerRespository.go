package repositories

import "github.com/WileESpaghetti/curt-db-utils/models"

type CustomerUserRepository interface {
	GetByApiKey(apiKey string) (c models.CustomerUser, err error)
	GetByEmail(email string) (c models.CustomerUser, err error)
}
