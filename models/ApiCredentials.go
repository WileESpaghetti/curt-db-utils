package models

import "time"

type ApiCredentials struct {
	Key       string    `json:"key" xml:"key,attr"`
	Type      string    `json:"type" xml:"type,attr"`
	TypeId    string    `json:"typeID" xml:"typeID,attr"`
	DateAdded time.Time `json:"date_added" xml:"date_added,attr"`
}
