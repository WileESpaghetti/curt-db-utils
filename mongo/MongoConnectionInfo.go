package mongo

import (
	"gopkg.in/mgo.v2"
	"os"
	"strings"
	"time"
)

// Builds MongoConnectionInfo from environment variables
// databaseVarName string Name of the environment variable containing the database you want to connect to
// userVarName string Name of the environment variable containing the username that is needed to authenticate to the desired database
// passwordVarName string Name of the environment variable containing the password that is needed to authenticate to the desired database
// hostVarName string Name of the environment variable containing the hostname (--host) running the mongodb server
func NewMongoDialInfoFromEnvironment(databaseVarName string, userVarName string, passwordVarName string, hostVarName string) *mgo.DialInfo {
	var info mgo.DialInfo

	info.Username = os.Getenv(userVarName)
	info.Password = os.Getenv(passwordVarName)
	info.Database = os.Getenv(databaseVarName)

	host := os.Getenv(hostVarName)
	if host == "" {
		host = "127.0.0.1"
	}
	availableHosts := strings.Split(host, ",")
	info.Addrs = append(info.Addrs, availableHosts...)

	// TODO move this into some sort of connection defaults
	info.Timeout = time.Second * 2
	info.FailFast = true
	info.Source = "admin"

	return &info
}
