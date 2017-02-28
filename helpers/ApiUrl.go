package helpers

import (
	"net/url"
	"database/sql/driver"
	"errors"
)

// Wrapper around url.URL so we can avoid manually having to use Parse() or String() when working with SQL databases
type ApiUrl struct {
	*url.URL
}

// Convert a Parsed URL back to a string
func (apiUrl ApiUrl) Value() (driver.Value, error) {
	return apiUrl.String(), nil
}

// Scan the URL string stored into the database into an ApiUrl
func (apiUrl *ApiUrl) Scan(value interface{}) error {
	if apiUrlVal, err := driver.String.ConvertValue(value); err == nil {
		if apiUrlStr, ok := apiUrlVal.(string); ok {
			embedded, err := url.Parse(apiUrlStr)
			*apiUrl = ApiUrl{embedded}
			return err
		}
	}
	return errors.New("failed to scan URL")
}

func NewApiUrl(str string) *ApiUrl {
	tmpUrl, _ := url.Parse(str)
	return &ApiUrl{tmpUrl}
}

