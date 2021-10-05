package database

import (
	"fmt"

	"example.com/fiber-apiserver/app/queries"
	"github.com/jmoiron/sqlx"
)

// Queries struct for collect all app queries.
type Queries struct {
	*queries.UserQueries // load queries from User model
	*queries.BookQueries // load queries from Book model
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection(n string) (*Queries, error) {

	var db *sqlx.DB
	var err error

	// Switch given names.
	switch n {
	case "postgres":
		// Define a new PostgreSQL connection.
		db, err = PostgreSQLConnection()
	case "mysql":
		// Define a new MySQL connection.
		db, err = MySQLConnection()
	default:
		// Return error message.
		return nil, fmt.Errorf("db name '%v' is not supported", n)
	}

	if err != nil {
		return nil, err
	}

	return &Queries{
		// Set queries from models:
		UserQueries: &queries.UserQueries{DB: db}, // from User model
		BookQueries: &queries.BookQueries{DB: db}, // from Book model
	}, nil
}
