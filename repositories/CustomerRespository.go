package repositories

import "github.com/WileESpaghetti/curt-db-utils/models"

type CustomerUserRepository interface {
	getByApiKey(apiKey string) models.Customer
	getByEmail(email string) models.Customer
}
