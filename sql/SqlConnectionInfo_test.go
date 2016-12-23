package sql

import (
	"testing"
	"os"
	"fmt"
)

func TestNewMysqlConnectionStringFromEnvironment(t *testing.T) {
	var actualConnectionString string
	var expectedConnectionString string

	hostKey  := "CURT_DB_UTILS_TEST_HOST"
	hostVal  := "test_hostname"

	userKey  := "CURT_DB_UTILS_TEST_USER"
	userVal  := "test_username"

	passKey  := "CURT_DB_UTILS_TEST_PASSWORD"
	passVal  := "test_password"

	dbKey    := "CURT_DB_UTILS_TEST_DATABASE"
	dbVal    := "test_database"

	protoKey := "CURT_DB_UTILS_TEST_PROTO"
	protoVal := "tcp"

	os.Setenv(hostKey,  hostVal)
	os.Setenv(userKey,  userVal)
	os.Setenv(passKey,  passVal)
	os.Setenv(dbKey,    dbVal)
	os.Setenv(protoKey, protoVal)

	expectedConnectionString =fmt.Sprintf("%s:%s@%s(%s)/%s?parseTime=%s&loc=%s", userVal, passVal, protoVal, hostVal, dbVal, "true", "America%2FChicago")
	actualConnectionString = NewMysqlConnectionStringFromEnvironment(dbKey, userKey, passKey, hostKey, protoKey)

	if (expectedConnectionString != actualConnectionString) {
		t.Error("Expected Connection String to be\n expected: %s\nactual:  %s", expectedConnectionString, actualConnectionString)
	}

	os.Unsetenv(hostKey)
	os.Unsetenv(userKey)
	os.Unsetenv(passKey)
	os.Unsetenv(dbKey)
	os.Unsetenv(protoKey)
}

func TestNewMysqlConnectionStringFromValues(t *testing.T) {
	var expectedConnectionString string
	var actualConnectionString string

	hostVal  := "test_hostname"
	userVal  := "test_username"
	passVal  := "test_password"
	dbVal    := "test_database"
	protoVal := "tcp"

	expectedConnectionString =fmt.Sprintf("%s:%s@%s(%s)/%s?parseTime=%s&loc=%s", userVal, passVal, protoVal, hostVal, dbVal, "true", "America%2FChicago")
	actualConnectionString = NewMysqlConnectionStringFromValues(dbVal, userVal, passVal, hostVal, protoVal)

	if (expectedConnectionString != actualConnectionString) {
		t.Error(fmt.Sprintf("Expected Connection String to be\nexpected: %s\nactual:   %s", expectedConnectionString, actualConnectionString))
	}
}

