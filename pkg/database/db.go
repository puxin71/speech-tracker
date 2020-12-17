package database

type DB interface {
	// Return all the talks and its attendant information from the database
	GetAllTalks() ([]Talk, error)
}
