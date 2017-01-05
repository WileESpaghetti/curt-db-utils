package repositories

import (
	"database/sql"
	"testing"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func TestGetByEmail(t *testing.T) {
	var err error
	testDb := "curt_db_utils_test"
	customerUserTable := "CustomerUser"
	session, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/")
	if err != nil {
		fmt.Println(err)
		t.Error("Could not connect to test database server")
		return
	}
	defer session.Close()

	_,err = session.Exec("DROP DATABASE IF EXISTS " + testDb)
	if err != nil {
		t.Error(err)
	}

	_,err = session.Exec("CREATE DATABASE " + testDb)
	if err != nil {
		panic(err)
	}

	_,err = session.Exec("USE " + testDb)
	if err != nil {
		panic(err)
	}

	_,err = session.Exec("CREATE TABLE " + customerUserTable + "(id integer, email varchar(255))")
	if err != nil {
		panic(err)
	}

	_,err = session.Exec("INSERT INTO " + customerUserTable + " (id, email) VALUES (1, 'example@example.com')")
	if err != nil {
		panic(err)
	}

	repo := SqlCustomerRepository{Session: session}
	customerUser, err := repo.GetByEmail("example@example.com")
	if ("1" != customerUser.Id) {
		t.Error("Expected CustomerUser.Id to be\n expected: %s\nactual:  %s", "1", customerUser.Id)
	}

	_,err = session.Exec("DROP DATABASE " + testDb)
	if err != nil {
		panic(err)
	}
}

func TestGetByApiKey(t *testing.T) {
	var err error
	testDb := "curt_db_utils_test"
	customerUserTable := "CustomerUser"
	apiKeyTable := "ApiKey"
	apiKeyTypeTable := "ApiKeyType"
	session, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/")
	if err != nil {
		fmt.Println(err)
		t.Error("Could not connect to test database server")
		return
	}
	defer session.Close()

	_,err = session.Exec("DROP DATABASE IF EXISTS " + testDb)
	if err != nil {
		t.Error(err)
	}

	_,err = session.Exec("CREATE DATABASE " + testDb)
	if err != nil {
		panic(err)
	}

	_,err = session.Exec("USE " + testDb)
	if err != nil {
		panic(err)
	}

	_,err = session.Exec("CREATE TABLE " + customerUserTable + "(id int, email varchar(255))")
	if err != nil {
		panic(err)
	}

	_,err = session.Exec("INSERT INTO " + customerUserTable + " (id, email) VALUES ('1', 'example@example.com')")
	if err != nil {
		panic(err)
	}

	// Add test API key
	testCustomerUserId := "1"
	testApiKeyTypeType := "TESTING_KEY"
	testApiKeyTypeId := "1"
	testApiKeyApiKey := "VALID_TEST_KEY"
	testInvalidApiKeyApiKey := "INVALID_TEST_KEY"
	testApiKeyId := "1"

	createApiKeyTypeTable := fmt.Sprintf("CREATE TABLE %s (id integer, type varchar(255))", apiKeyTypeTable)
	_,err = session.Exec(createApiKeyTypeTable)
	if err != nil {
		panic(err)
	}

	insertTestApiKeyType := fmt.Sprintf("INSERT INTO %s (id, type) VALUES (%s, '%s')", apiKeyTypeTable, testApiKeyTypeId, testApiKeyTypeType)
	_,err = session.Exec(insertTestApiKeyType)
	if err != nil {
		panic(err)
	}

	createApiKeyTable := fmt.Sprintf("CREATE TABLE %s (id integer, user_id integer, type_id integer, api_key varchar(255))", apiKeyTable)
	_,err = session.Exec(createApiKeyTable)
	if err != nil {
		panic(err)
	}

	insertTestApiKey := fmt.Sprintf("INSERT INTO %s (id, user_id, type_id, api_key) VALUES (%s, %s, %s, '%s')",
		apiKeyTable, testApiKeyId, testCustomerUserId, testApiKeyTypeId, testApiKeyApiKey)
	_,err = session.Exec(insertTestApiKey)
	if err != nil {
		panic(err)
	}

	// Test the result
	repo := SqlCustomerRepository{Session: session}
	customerUser, err := repo.GetByApiKey(testApiKeyApiKey, testApiKeyTypeType)
	if (fmt.Sprintf(customerUser.Id) != testCustomerUserId) {
		t.Errorf("Expected CustomerUser.Id to be\n expected: %s\nactual:  %s", testCustomerUserId, customerUser.Id)
	}
	if (err != nil) {
		t.Error("Unexpected error occured from GetUserByApiKey")
	}

	_, err = repo.GetByApiKey(testInvalidApiKeyApiKey, testApiKeyTypeType)
	if (err == nil) {
		t.Error("Error not thrown for invalid API key")
	}

	_,err = session.Exec("DROP DATABASE " + testDb)
	if err != nil {
		panic(err)
	}
}
