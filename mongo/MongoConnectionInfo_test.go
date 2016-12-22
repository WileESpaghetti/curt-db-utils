package mongo

import (
	"testing"
	"os"
	"gopkg.in/mgo.v2"
	"strings"
)

func TestNewMongoDialInfoFromEnvironment(t *testing.T) {
	var dialInfo *mgo.DialInfo

	hostKey := "CURT_DB_UTILS_TEST_HOST"
	hostVal := "test_hostname"

	userKey := "CURT_DB_UTILS_TEST_USER"
	userVal := "test_username"

	passKey := "CURT_DB_UTILS_TEST_PASSWORD"
	passVal := "test_password"

	dbKey   := "CURT_DB_UTILS_TEST_DATABASE"
	dbVal   := "test_database"

	os.Setenv(hostKey, hostVal)
	os.Setenv(userKey, userVal)
	os.Setenv(passKey, passVal)
	os.Setenv(dbKey,   dbVal)

	dialInfo = NewMongoDialInfoFromEnvironment(dbKey, userKey, passKey, hostKey)

	if dialInfo.Database != dbVal {
		t.Error("Expected Database to be %s, got %s", dbVal, dialInfo.Database)
	}

	if dialInfo.Username != userVal {
		t.Error("Expected Username to be %s, got %s", userVal, dialInfo.Username)
	}

	if dialInfo.Password != passVal {
		t.Error("Expected Password to be %s, got %s", passVal, dialInfo.Password)
	}

	expectedHosts := strings.Split(hostVal, ",")
	for index := range expectedHosts {
		if dialInfo.Addrs[index] != expectedHosts[index] {
			t.Error("Expected Addrs to contain %q, got %q", hostVal, dialInfo.Addrs)
		}
	}

	os.Unsetenv(hostKey)
	os.Unsetenv(userKey)
	os.Unsetenv(passKey)
	os.Unsetenv(dbKey)
}
