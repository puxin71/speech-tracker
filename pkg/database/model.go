package database

import (
	"fmt"
	"time"
)

// RoleType identifies if the attendant is a speaker or an attendee in a talk
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
	Title string `json:"title,omitempty"`
	// Abstract description of the talk
	Abstract string `json:"abstract,omitempty"`
	// Room booked for the talk
	Room int `json:"room,omitempty"`
	// Speaker who presents in the talk
	Speaker Attendant `json:"speaker,omitempty"`
	// Attendees that are registered to the talk
	Attendees []Attendant `json:"attendees,omitempty"`
}

type Attendant struct {
	// First name and last name of the attendant
	Name string `json:"name,omitempty"`
	// Company of the attendant
	Company string `json:"company,omitempty"`
	// Email contact of the attendant
	Email string `json:"email,omitempty"`
	// Attendees' registeration time in UTC
	Registered time.Time `json:"registered,omitempty"`
	// Identifies if the attendant is a speaker or an attendee
	// This field is ignored in the JSON output
	Role RoleType `json:"-"`
	// A short biography of the attendant
	Bio string `json:"bio,omitempty"`
}
