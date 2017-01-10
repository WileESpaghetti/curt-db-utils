package models

import (
	"database/sql"
	"fmt"
	"net/url"
)

type SqlBrandRepository struct {
	Session *sql.DB
}

func (repo SqlBrandRepository) GetById(id int) (brand Brand, err error) {
	var logo *string
	var logoAlt *string

	getBrandById := `select ID, name, code, logo, logoAlt, formalName, longName, primaryColor, autocareID
			from Brand where ID = ? limit 1`

	stmt, err := repo.Session.Prepare(getBrandById)
	if err != nil {
		return brand, err
	}
	defer stmt.Close()

	result := stmt.QueryRow(id)
	err = result.Scan(&brand.ID, &brand.Name, &brand.Code, &logo, &logoAlt, &brand.FormalName, &brand.LongName, &brand.PrimaryColor, &brand.AutocareID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("%s", "brand doesn't exist")
		}
		return brand, err
	}

	if logo != nil {
		brand.Logo, _ = url.Parse(*logo)
	}

	if logoAlt != nil {
		brand.LogoAlternate, _ = url.Parse(*logoAlt)
	}

	return brand, err
}

func (repo SqlBrandRepository) SaveNew(brand Brand) (err error) {
	saveNewBrand := `insert into Brand(name, code, logo, logoAlt, formalName, longName, primaryColor, autocareID) values (?,?,?,?,?,?,?,?)`
	stmt, err := repo.Session.Prepare(saveNewBrand)
	if err != nil {
		return err
	}
	defer stmt.Close()

	logo := brand.Logo.String()
	logoAlt := brand.Logo.String()
	_, err = stmt.Exec(&brand.Name, &brand.Code, &logo, &logoAlt, &brand.FormalName, &brand.LongName, &brand.PrimaryColor, &brand.AutocareID)
	if err != nil {
		return err
	}

	//id, err := result.LastInsertId()
	//if err != nil {
	//	return err
	//}
	//brand.ID = int(id)

	return err
}
