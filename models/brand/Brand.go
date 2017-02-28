package brand

import dbUtils "github.com/WileESpaghetti/curt-db-utils/helpers"

type Brand struct {
	ID            int
	Name          string
	Code          string
	Logo          *dbUtils.ApiUrl
	LogoAlternate *dbUtils.ApiUrl
	FormalName    string
	LongName      string
	PrimaryColor  string
	AutocareID    string
}
