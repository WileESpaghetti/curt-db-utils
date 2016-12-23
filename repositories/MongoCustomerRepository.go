package repositories

import (
	"github.com/WileESpaghetti/curt-db-utils/models"
	"gopkg.in/mgo.v2"
)

type MongoCustomerRepository struct {
	session *mgo.Session
}

func (repo MongoCustomerRepository) getByApiKey(apiKey string) models.Customer {
	// FIXME not implemented
	return nil
}

func (repo MongoCustomerRepository) getByEmail(apiKey email) models.Customer {
	// FIXME not implemented
	return nil
}
