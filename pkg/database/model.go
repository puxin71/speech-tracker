package database

import (
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
