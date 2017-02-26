package brand

import "net/url"

type Brand struct {
	ID            int
	Name          string
	Code          string
	Logo          *url.URL
	LogoAlternate *url.URL
	FormalName    string
	LongName      string
	PrimaryColor  string
	AutocareID    string
}
