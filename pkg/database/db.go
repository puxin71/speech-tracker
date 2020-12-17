package database

type DB interface {
	// Return all the talks and its attendant information from the database
	GetAll() ([]Talk, error)
}
