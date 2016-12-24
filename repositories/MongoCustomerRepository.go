package repositories

import (
	"github.com/WileESpaghetti/curt-db-utils/models"
	"gopkg.in/mgo.v2"
)

type MongoCustomerUserRepository struct {
	session *mgo.Session
}

func (repo MongoCustomerUserRepository) getByApiKey(apiKey string) (c models.CustomerUser, err error) {
	// FIXME not implemented
	return models.CustomerUser{}, nil
}

func (repo MongoCustomerUserRepository) getByEmail(email string) (c models.CustomerUser, err error) {
	// FIXME not implemented
	return models.CustomerUser{},  nil
}
