package brand

import (
	"testing"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	dbUtils "github.com/WileESpaghetti/curt-db-utils/helpers"
)

func TestSqlBrandRepository_GetByIdShouldReturnBrandWhenGivenValidId(t *testing.T) {
	// Setup expected Result
	expectedBrand := Brand {ID: 1,
		Name: "ExampleName",
		Code: "ExampleCode",
		Logo: dbUtils.NewApiUrl("http://www.example.com/logo.png"),
		LogoAlternate: dbUtils.NewApiUrl("http://www.example.com/logo_alt.png"),
		FormalName: "Example Formal Name",
		LongName: "Example Long Name",
		PrimaryColor: "#ffffff",
		AutocareID: "EXAM"}

	// Setup the database mock
	session, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer session.Close()

	rows := sqlmock.NewRows([]string{"ID", "name", "code", "logo", "logoAlt", "formalName", "longName",
		"primaryColor", "autocareID"})
	rows.AddRow(expectedBrand.ID, expectedBrand.Name, expectedBrand.Code, expectedBrand.Logo, expectedBrand.LogoAlternate, expectedBrand.FormalName,
		expectedBrand.LongName, expectedBrand.PrimaryColor, expectedBrand.AutocareID)

	mock.ExpectPrepare("SELECT (.+)").ExpectQuery().WillReturnRows(rows)

	// Test SqlBrandRepository.GetById
	repo := SqlBrandRepository{Session: session}
	brand, err := repo.GetById(expectedBrand.ID)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when looking up Brand by ID", err)
	}

	if ! reflect.DeepEqual(brand, expectedBrand) {
		t.Errorf("Actual Brand Does not match expected\n  Actual: %+v\nExpected: %+v", brand, expectedBrand)
	}
}

/*
func TestSqlBrandRepository_SaveNew(t *testing.T) {
	// Setup expected Result
	expectedLogo := "http://www.example.com/logo.png"
	expectedLogoAlternate := "http://www.example.com/logo_alt.png"
	expectedLogoParsed, _ := url.Parse(expectedLogo)
	expectedLogoAlternateParsed ,_ := url.Parse(expectedLogoAlternate)
	expectedBrand := Brand {ID: 1,
		Name: "ExampleName",
		Code: "ExampleCode",
		Logo: expectedLogoParsed,
		LogoAlternate: expectedLogoAlternateParsed,
		FormalName: "Example Formal Name",
		LongName: "Example Long Name",
		PrimaryColor: "#ffffff",
		AutocareID: "EXAM"}

	// Setup the database mock
	session, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer session.Close()


	fmt.Printf("%+v\n", mock.ExpectPrepare("INSERT INTO Brand.+").ExpectExec().
		WillReturnResult(sqlmock.NewResult(0, 1)))

	// Test SqlBrandRepository.GetById
	repo := SqlBrandRepository{Session: session}
	err = repo.SaveNew(expectedBrand)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}
*/

// AS A ???
// I WANT to save a Brand to the database
// SO THAT I can retrieve it later
/*
TODO convert to use ApiURL
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
 */
