package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	Username = "root"
	Password = "root"
	Hostname = "172.29.0.2:3306"
	DbName   = "test_db"
)

// MySQL database client
type MySQLDB struct {
	Querier
	Upserter
	db *sql.DB
}

// Return the database source name
func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", Username, Password, Hostname, DbName)
}

// Connects to MySQL database with a connection pool
func NewMySQLDB() (MySQLDB, error) {
	db, err := sql.Open("mysql", dsn(DbName))
	if err != nil {
		log.Fatalln("fail to connect to MySQL server", err)
	}

	return MySQLDB{db: db}, nil
}

// Close connection pool to the MySQL database
func (db MySQLDB) Close() {
	db.db.Close()
}

// Check if the connection to MySQL server is alive
func (db MySQLDB) Ping(ctx context.Context, timeout time.Duration) error {
	var err error
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	status := "up"
	if err = db.db.PingContext(ctx); err != nil {
		status = "down"
	}
	log.Println("MySQL server is", status)
	return err
}
