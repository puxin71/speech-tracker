package database

import (
	"encoding/json"
	"fmt"
	"time"
)

// RoleType identifies if the attendant is a speaker or an attendant to a talk
type RoleType int

const (
	// Default role is an attendee
	ATTENDEE RoleType = 0
	// Otherwise it is a speaker
	SPEAKER RoleType = 1
)

func (e RoleType) String() string {
	switch e {
	case SPEAKER:
		return "speaker"
	case ATTENDEE:
		return "attendee"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

type Talk struct {
	// Title of the talk
	Title string `json:"title"`
	// Abstract description of the talk
	Abstract string `json:"abstract"`
	// Room booked for the talk
	Room int `json:"room"`
	// Speaker who presents in the talk
	Speaker Attendant `json:"speaker"`
	// Attendees that are registered to the talk
	Attendees []Attendant `json:"attendees"`
}

type Attendant struct {
	// First name and last name of the attendant
	Name string `json:"name"`
	// Company of the attendant
	Company string `json:"company"`
	// Email contact of the attendant
	Email string `json:"email"`
	// UTC time when the attendant registers to the talk
	Registered time.Time `json:"registered"`
	// Identifies if the attendant is a speaker or an attendee
	Role RoleType
	// A short biography of the attendant
	Bio string `json:"bio"`
}

// Custom Json Marshal to ignore the Role field
func (a Attendant) MarshalJSON() ([]byte, error) {
	var tmp struct {
		Name       string `json:"name"`
		Company    string `json:"company"`
		Email      string `json:"email"`
		Registered string `json:"registered"`
		Bio        string `json:"bio"`
	}

	tmp.Name = a.Name
	tmp.Company = a.Company
	tmp.Email = a.Email
	tmp.Registered = format(a.Registered)
	tmp.Bio = a.Bio
	return json.Marshal(&tmp)
}

// Always display the time with the RFC3339 format
func format(t time.Time) string {
	return t.Format(time.RFC3339)
}
