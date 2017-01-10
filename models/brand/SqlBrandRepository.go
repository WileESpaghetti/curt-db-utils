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
	getBrandById := `select ID, name, code, logo, logoAlt, formalName, longName, primaryColor, autocareID
			from Brand where ID = ? limit 1`
	//var b Brand
	//var err error

	stmt, err := repo.Session.Prepare(getBrandById)
	if err != nil {
		return brand, err
	}
	defer stmt.Close()

	var logo, logoAlt, formal, long, primary, autocare *string
	result := stmt.QueryRow(id)
	err = result.Scan(&brand.ID, &brand.Name, &brand.Code, &logo, &logoAlt, &formal, &long, &primary, &autocare)
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

	if formal != nil {
		brand.FormalName = *formal
	}

	if long != nil {
		brand.LongName = *long
	}

	if primary != nil {
		brand.PrimaryColor = *primary
	}

	if autocare != nil {
		brand.AutocareID = *autocare
	}

	return brand, err
}
