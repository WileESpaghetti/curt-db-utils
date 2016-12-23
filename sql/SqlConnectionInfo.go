package sql

import (
	"fmt"
	"os"
)

// Creates a string formatted in a way that the (go-sql-driver)[https://github.com/go-sql-driver/mysql] can connect to a MySQL instance
func NewMysqlConnectionStringFromEnvironment(databaseVarName string, userVarName string, passwordVarName string, hostVarName string, protocolVarName string) string {
	// QUESTION does any of this need to be URL encoded?
	user  := os.Getenv(userVarName)
	pass  := os.Getenv(passwordVarName)
	db    := os.Getenv(databaseVarName)
	proto := os.Getenv(protocolVarName)
	host  := os.Getenv(hostVarName)

	// TODO move into some sort of defaults struct
	time  := "America%2FChicago"
	parse := "true"

	return fmt.Sprintf("%s:%s@%s(%s)/%s?parseTime=%s&loc=%s", user, pass, proto, host, db, parse, time)
}
