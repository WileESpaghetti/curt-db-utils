package brand

import (
	"database/sql"
	"testing"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
	"reflect"
)

func TestGetById(t *testing.T) {
	// FIXME replace individual tests by creating Brand{} instance and deepEqual to resulting Brand
	var err error
	testDb := "curt_db_utils_test"
	brandTable := "Brand"
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

	_,err = session.Exec("CREATE TABLE " + brandTable + "(`ID` int(11) NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, `code` varchar(255) NOT NULL, `logo` varchar(255) DEFAULT NULL, `logoAlt` varchar(255) DEFAULT NULL, `formalName` varchar(255) DEFAULT NULL, `longName` varchar(255) DEFAULT NULL, `primaryColor` varchar(10) DEFAULT NULL, `autocareID` varchar(4) DEFAULT NULL, PRIMARY KEY (`ID`))")
	if err != nil {
		panic(err)
	}

	testId := 1
	testName := "TestBrandName"
	testCode := "TestCode"
	testLogo := "http://www.example.com/logo.png"
	testLogoAlt := "http://www.example.org/logo-alt.png"
	testFormalName := "TestFormalName"
	testLongName := "Test Long Name"
	testPrimaryColor := "red"
	testAutoCareID := "test"

	insertTestBrand := fmt.Sprintf("INSERT INTO %s (ID, name, code, logo, logoAlt, formalName, longName, primaryColor, autocareID) VALUES (%d, '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')",
		brandTable,
		testId,
		testName,
		testCode,
		testLogo,
		testLogoAlt,
		testFormalName,
		testLongName,
		testPrimaryColor,
		testAutoCareID)
	_,err = session.Exec(insertTestBrand)
	if err != nil {
		panic(err)
	}

	repo := SqlBrandRepository{Session: session}
	brand, err := repo.GetById(testId)
	if (testId != brand.ID) {
		t.Errorf("Expected Brand.ID to be\n expected: %s\nactual:  %s", testId, brand.ID)
	}

	if (testName != brand.Name) {
		t.Errorf("Expected Brand.Name to be\n expected: %s\nactual:  %s", testName, brand.Name)
	}

	if (testCode != brand.Code) {
		t.Errorf("Expected Brand.Code to be\n expected: %s\nactual:  %s", testCode, brand.Code)
	}

	parsedLogo, _ := url.Parse(testLogo)
	if (!reflect.DeepEqual(parsedLogo, brand.Logo)) {
		t.Errorf("Expected Brand.Logo to be\n expected: %s\nactual:  %s", parsedLogo, brand.Logo)
	}

	parsedLogoAlt, _ := url.Parse(testLogoAlt)
	if (!reflect.DeepEqual(parsedLogoAlt, brand.LogoAlternate)) {
		t.Errorf("Expected Brand.Logo to be\n expected: %s\nactual:  %s", parsedLogo, brand.Logo)
	}

	if (testFormalName != brand.FormalName) {
		t.Errorf("Expected Brand.FormalName to be\n expected: %s\nactual:  %s", testFormalName, brand.FormalName)
	}

	if (testPrimaryColor != brand.PrimaryColor) {
		t.Errorf("Expected Brand.PrimaryColor to be\n expected: %s\nactual:  %s", testPrimaryColor, brand.PrimaryColor)
	}

	if (testAutoCareID != brand.AutocareID) {
		t.Errorf("Expected Brand.AutoCareAutoCareID to be\n expected: %s\nactual:  %s", testAutoCareID, brand.AutocareID)
	}

	_,err = session.Exec("DROP DATABASE " + testDb)
	if err != nil {
		panic(err)
	}
}

// AS A ???
// I WANT to save a Brand to the database
// SO THAT I can retrieve it later
func TestSaveNew(t *testing.T) {
	var err error
	var originalCount, expectedCount, actualCount int

	testDb := "curt_db_utils_test"
	brandTable := "Brand"
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

	_,err = session.Exec("CREATE TABLE " + brandTable + "(`ID` int(11) NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, `code` varchar(255) NOT NULL, `logo` varchar(255) DEFAULT NULL, `logoAlt` varchar(255) DEFAULT NULL, `formalName` varchar(255) DEFAULT NULL, `longName` varchar(255) DEFAULT NULL, `primaryColor` varchar(10) DEFAULT NULL, `autocareID` varchar(4) DEFAULT NULL, PRIMARY KEY (`ID`))")
	if err != nil {
		panic(err)
	}

	testName := "TestBrandName"
	testCode := "TestCode"
	testLogo := "http://www.example.com/logo.png"
	parsedLogo, err := url.Parse(testLogo)
	testLogoAlt := "http://www.example.org/logo-alt.png"
	parsedLogoAlt, err := url.Parse(testLogoAlt)
	testFormalName := "TestFormalName"
	testLongName := "Test Long Name"
	testPrimaryColor := "red"
	testAutoCareID := "test"

	testBrand := Brand{
		Name: testName,
		Code: testCode,
		Logo: parsedLogo,
		LogoAlternate: parsedLogoAlt,
		FormalName: testFormalName,
		LongName: testLongName,
		PrimaryColor: testPrimaryColor,
		AutocareID: testAutoCareID }


	countSavedBrands := fmt.Sprintf("SELECT COUNT(*) from %s", brandTable)

	originalBrandCount, err := session.Query(countSavedBrands)
	if (err != nil) {
		t.Error("Unexpected error retreiving the original Brand count")
		t.Error(err)
	}
	originalBrandCount.Next()
	originalBrandCount.Scan(&originalCount)
	expectedCount = originalCount + 1
	originalBrandCount.Close()

	// test begins here
	repo := SqlBrandRepository{Session: session}
	err = repo.SaveNew(testBrand)

	actualBrandCount, err := session.Query(countSavedBrands)
	if (err != nil) {
		t.Error("Unexpected error retreiving the new Brand count")
		t.Error(err)
	}
	actualBrandCount.Next()
	actualBrandCount.Scan(&actualCount)
	originalBrandCount.Close()

	if expectedCount != actualCount {
		t.Errorf("Expected Brand count did not match\n expected: %d\nactual:  %d", expectedCount, actualCount)
	}


	_,err = session.Exec("DROP DATABASE " + testDb)
	if err != nil {
		panic(err)
	}
}
