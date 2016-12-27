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

	repo := SqlCustomerRepository{session: session}
	customerUser, err := repo.GetByEmail("example@example.com")
	if ("1" != customerUser.Id) {
		t.Error("Expected CustomerUser.Id to be\n expected: %s\nactual:  %s", "1", customerUser.Id)
	}

	_,err = session.Exec("DROP DATABASE " + testDb)
	if err != nil {
		panic(err)
	}
}

