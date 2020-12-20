package database

import "database/sql"

// Provides the APIs to query, insert, and upsert rows to a RDBMS database, i,e MySQL
type SQLStore struct {
	Querier
	Upserter
	db *sql.DB
}

func NewStore(db *sql.DB) SQLStore {
	return SQLStore{db: db}
}
