package repositories

import "github.com/WileESpaghetti/curt-db-utils/models"

type CustomerUserRepository interface {
	getByApiKey(apiKey string) models.CustomerUser
	getByEmail(email string) models.CustomerUser
}
