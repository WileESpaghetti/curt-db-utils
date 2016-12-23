package repositories

import (
	"github.com/WileESpaghetti/curt-db-utils/models"
	"gopkg.in/mgo.v2"
)

type MongoCustomerUserRepository struct {
	session *mgo.Session
}

func (repo MongoCustomerUserRepository) getByApiKey(apiKey string) models.Customer {
	// FIXME not implemented
	return nil
}

func (repo MongoCustomerUserRepository) getByEmail(apiKey email) models.Customer {
	// FIXME not implemented
	return nil
}
