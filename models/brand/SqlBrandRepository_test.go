package brand

import (
	"testing"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	dbUtils "github.com/WileESpaghetti/curt-db-utils/helpers"
	"strings"
	"database/sql/driver"
)

func TestSqlBrandRepository_WhenGivenAValidI_GetById_ShouldReturnAStoredBrand(t *testing.T) {
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

func TestSqlBrandRepository_WhenGivenAnInvalidI_GetById_ShouldReturnABrandNotFoundError(t *testing.T) {
	// Setup the database mock
	session, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer session.Close()

	rows := sqlmock.NewRows([]string{"ID", "name", "code", "logo", "logoAlt", "formalName", "longName",
		"primaryColor", "autocareID"})

	mock.ExpectPrepare("SELECT (.+)").ExpectQuery().WillReturnRows(rows)

	// Test SqlBrandRepository.GetById
	invalidId := -1
	repo := SqlBrandRepository{Session: session}
	_, err = repo.GetById(invalidId)
	if strings.Index(err.Error(), BRAND_NOT_FOUND) == -1 {
		t.Fatalf("GetById() did not return a BRAND_NOT_FOUND error when looking up an invalid ID\n" +
			"Got this instead: %s", err)
	}
}

func TestSqlBrandRepository_WhenGivenAnErrorPreparing_GetById_ShouldReturnAnError(t *testing.T) {
	// Setup the database mock
	session, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer session.Close()

	mock.ExpectPrepare("SELECT (.+)").WillReturnError(driver.ErrBadConn) // amoung other errors

	// Test SqlBrandRepository.GetById
	arbitraryId := 1
	repo := SqlBrandRepository{Session: session}
	_, err = repo.GetById(arbitraryId)
	if err == nil {
		t.Fatal("GetById() did not return an error from a bad Prepare()\n")
	}
}

func TestSqlBrandRepository_SaveNew(t *testing.T) {
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


	mock.ExpectPrepare("INSERT INTO Brand.+").ExpectExec().
		WithArgs(expectedBrand.Name,
			expectedBrand.Code,
			expectedBrand.Logo.String(),
			expectedBrand.LogoAlternate.String(),
			expectedBrand.FormalName,
			expectedBrand.LongName,
			expectedBrand.PrimaryColor,
			expectedBrand.AutocareID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	// Test SqlBrandRepository.GetById
	repo := SqlBrandRepository{Session: session}
	err = repo.SaveNew(expectedBrand)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestSqlBrandRepository_WhenGivenAnErrorPreparing_SaveNew_ShouldReturnAnError(t *testing.T) {
	// Setup the database mock
	session, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer session.Close()

	mock.ExpectPrepare("INSERT INTO Brand.+").WillReturnError(driver.ErrBadConn) // amoung other errors

	// Test SqlBrandRepository.GetById
	repo := SqlBrandRepository{Session: session}
	err = repo.SaveNew(Brand{})
	if err == nil {
		t.Fatal("GetById() did not return an error from a bad Prepare()\n")
	}
}

// AS A ???
// I WANT to save a Brand to the database
// SO THAT I can retrieve it later
