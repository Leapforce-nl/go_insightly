package insightly

import (
	"strings"
	"time"
)

type Contact struct {
	CONTACT_ID       int           `json:"CONTACT_ID"`
	FIRST_NAME       string        `json:"FIRST_NAME"`
	LAST_NAME        string        `json:"LAST_NAME"`
	ORGANISATION_ID  int           `json:"ORGANISATION_ID"`
	EMAIL_ADDRESS    string        `json:"EMAIL_ADDRESS"`
	DATE_UPDATED_UTC string        `json:"DATE_UPDATED_UTC"`
	CUSTOMFIELDS     []CustomField `json:"CUSTOMFIELDS"`
	DateUpdated      time.Time
	Initials         string
	Gender           string
	Title            string
	Email            string
	IsMainContact    bool
	PushToEO         bool
}

/*
type iContacts struct {
	Contacts []Contact
}*/

func (c *Contact) iGenderToGender(gender string) string {
	if strings.ToLower(gender) == "man" {
		return "Mannelijk"
	} else if strings.ToLower(gender) == "vrouw" {
		return "Vrouwelijk"
	}

	return "Onbekend"
}

func (c *Contact) iGenderToTitle(gender string) string {
	if strings.ToLower(gender) == "man" {
		return "DHR"
	} else if strings.ToLower(gender) == "vrouw" {
		return "MEVR"
	}

	return ""
}