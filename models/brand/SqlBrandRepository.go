package brand

import (
	"database/sql"
	"github.com/pkg/errors"
)

type SqlBrandRepository struct {
	Session *sql.DB
}

const BRAND_NOT_FOUND = "Brand doesn't exist"
const BRAND_CONNECTION_ISSUE = ""

func (repo SqlBrandRepository) GetById(id int) (brand Brand, err error) {
	getBrandById := `SELECT ID, name, code, logo, logoAlt, formalName, longName, primaryColor, autocareID
			FROM Brand where ID = ?`

	stmt, err := repo.Session.Prepare(getBrandById)
	if err != nil {
		return brand, errors.Wrap(err, BRAND_CONNECTION_ISSUE)
	}
	defer stmt.Close()

	result := stmt.QueryRow(id)
	err = result.Scan(&brand.ID, &brand.Name, &brand.Code, &brand.Logo, &brand.LogoAlternate, &brand.FormalName, &brand.LongName, &brand.PrimaryColor, &brand.AutocareID)
	// FIXME I think this overrides the expected errors below
	switch {
	case err == sql.ErrNoRows:
		err = errors.Wrap(err, BRAND_NOT_FOUND)
		fallthrough
	case err != nil:
		return brand, err
	}

	return brand, err
}

func (repo SqlBrandRepository) SaveNew(brand Brand) (err error) {
	saveNewBrand := `INSERT INTO Brand(name, code, logo, logoAlt, formalName, longName, primaryColor, autocareID) values (?,?,?,?,?,?,?,?)`
	stmt, err := repo.Session.Prepare(saveNewBrand)
	if err != nil {
		return err
	}
	defer stmt.Close()

	logo := brand.Logo.String()
	logoAlt := brand.LogoAlternate.String()
	_, err = stmt.Exec(&brand.Name, &brand.Code, &logo, &logoAlt, &brand.FormalName, &brand.LongName, &brand.PrimaryColor, &brand.AutocareID)

	return err
}

// TODO should not leak mysql connections
