package repositories

import "github.com/WileESpaghetti/curt-db-utils/models"

type CustomerRepository interface {
	getByApiKey(apiKey string) models.Customer
}
