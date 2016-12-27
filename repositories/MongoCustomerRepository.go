package repositories

import (
	"github.com/WileESpaghetti/curt-db-utils/models"
	"gopkg.in/mgo.v2"
)

type MongoCustomerUserRepository struct {
	Session *mgo.Session
}

func (repo MongoCustomerUserRepository) GetByApiKey(apiKey string) (c models.CustomerUser, err error) {
	// FIXME not implemented
	return models.CustomerUser{}, nil
}

func (repo MongoCustomerUserRepository) GetByEmail(email string) (c models.CustomerUser, err error) {
	// FIXME not implemented
	return models.CustomerUser{},  nil
}
