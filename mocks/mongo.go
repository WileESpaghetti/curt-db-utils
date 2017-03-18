package mocks

import (
	"github.com/ory-am/dockertest"
	"gopkg.in/mgo.v2"
	"fmt"
	"github.com/pkg/errors"
)

// use the default pool for dockertest.
// @see [dockertest.NewPool()](https://github.com/ory/dockertest/blob/v3/dockertest.go#L63)
const DEFAULT_DOCKER_POOL = ""

// Docker container configuration
const (
	// Image that should be pulled down to build the docker container
	IMAGE = "mongo"

	// This should match what we are currently using in production
	TAG = "3.2"

	// Hostname of the running Docker container
	HOST = "localhost"

	// The name of the exposed port rather than the actual port number
	PORT = "27017/tcp"
)

// Docker related error messages
const (
	ERR_DOCKER_COULD_NOT_CONNECT = "Could not connect to docker: %s"
	ERR_DOCKER_RESOURCE_START = "Could not start resource: %s"
)

// Connection information for the running MongoDB container
type Mongo struct {
	Pool *dockertest.Pool
	Resource *dockertest.Resource
	Session *mgo.Session
}

func NewSessionFromResource(mgoPool *dockertest.Pool, mgoResource *dockertest.Resource) (session *mgo.Session, err error) {
	//err = mgoPool.Retry(attemptDialWith(session, mgoResource))
	err = mgoPool.Retry(func () (err error) {
		session, err = mgo.Dial(fmt.Sprintf("%s:%s", HOST, mgoResource.GetPort(PORT)))
		if err != nil {
			return err
		}

		return session.Ping()
	})
	return session, err
}

// Spin up a new Docker container running MongoDB
func NewDockertestMongo() (mock Mongo, err error) {
	mgoPool, err := dockertest.NewPool(DEFAULT_DOCKER_POOL)
	if err != nil {
		return mock, errors.Wrap(err, ERR_DOCKER_COULD_NOT_CONNECT)
	}

	resource, err := mgoPool.Run(IMAGE, TAG, nil)
	if err != nil {
		return mock, errors.Wrap(err, ERR_DOCKER_RESOURCE_START)
	}

	session, err := NewSessionFromResource(mgoPool, resource)
	if err != nil {
		return mock, errors.Wrap(err, ERR_DOCKER_COULD_NOT_CONNECT)
	}

	mock = Mongo{Pool: mgoPool,
		Resource: resource,
		Session: session}

	return mock, err
}

// Stop the running Docker container and clean up
func (mock Mongo) Close() (err error) {
	return mock.Pool.Purge(mock.Resource)
}
