package database

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"example.com/fiber-apiserver/pkg/utils"
	// _ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// MySQLConnection func for connection to MySQL database.
func MySQLConnection() (*sqlx.DB, error) {
	// Define database connection settings.
	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	// Build MySQL connection URL.
	mysqlConnURL, err := utils.ConnectionURLBuilder("mysql")
	if err != nil {
		return nil, err
	}

	// Define database connection for MySQL.
	// db, err = sql.Open("mysql", dbUser+":"+dbPassword+"@"+dbProtocol+"("+dbAddress+":"+dbPort+")/"+dbName)
	// db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")
	// db, err := sqlx.Connect("mysql", "test:test@(localhost:3306)/test")
	db, err := sqlx.Connect("mysql", mysqlConnURL)
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	// Set database connection settings:
	db.SetMaxOpenConns(maxConn)                           // the default is 0 (unlimited)
	db.SetMaxIdleConns(maxIdleConn)                       // defaultMaxIdleConns = 2
	db.SetConnMaxLifetime(time.Duration(maxLifetimeConn)) // 0, connections are reused forever

	// Try to ping database.
	if err := db.Ping(); err != nil {
		defer db.Close() // close database connection
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}

	return db, nil
}

// db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")
// if err != nil {
// 	panic(err)
// }

// See "Important settings" section.
// db.SetConnMaxLifetime(time.Minute * 3)
// db.SetMaxOpenConns(10)
// db.SetMaxIdleConns(10)

// fmt.Println("connect success", db)

// defer db.Close()
