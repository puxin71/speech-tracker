package database

type Querier interface {
	// Return all the talks and its attendant information from the database
	GetAllTalks() ([]Talk, error)

	// Return all the attendees that are registered for a talk
	GetAttendees(tkid int) ([]Attendant, error)
}

type Upserter interface {
}

type Store interface {
	Querier
	Upserter
}
