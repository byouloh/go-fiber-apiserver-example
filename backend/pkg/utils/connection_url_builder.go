package utils

import (
	"fmt"
	"os"
)

// ConnectionURLBuilder func for building URL connection.
func ConnectionURLBuilder(n string) (string, error) {
	// Define URL to connection.
	var url string

	// Switch given names.
	switch n {
	case "postgres":
		// URL for PostgreSQL connection.
		// "postgres://YourUserName:YourPassword@YourHostname:5432/YourDatabaseName"
		// DATABASE_URL=postgres://{user}:{password}@{hostname}:{port}/{database-name}
		// db, err := sql.Open("postgres", os.Getenv("DB_URL"))
		// DB_URL=postgres://testuser:testpassword@localhost/testmooc
		// db, err := sqlx.Connect("pgx","postgresql://user:pass@localhost:5433/mydb")
		// const (
		// 	host     = "localhost"
		// 	port     = 5432
		// 	user     = "postgres"
		// 	password = "your-password"
		// 	dbname   = "calhounio_demo"
		//   )
		// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		//     "password=%s dbname=%s sslmode=disable",
		//     host, port, user, password, dbname)
		//   db, err := sql.Open("postgres", psqlInfo)
		url = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_NAME"),
			os.Getenv("POSTGRES_SSL_MODE"),
		)
	case "mysql":
		// URL for MySQL connection.
		// db, err = sql.Open("mysql", dbUser+":"+dbPassword+"@"+dbProtocol+"("+dbAddress+":"+dbPort+")/"+dbName)
		// db, err := sqlx.Connect("mysql", "test:test@(localhost:3306)/test")
		// "user=%s:password=%s@host=%s:port=%s/dbname=%s", "%s:%s@%s:%s/%s"
		// DATABASE_URL=mysql://{user}:{password}@{hostname}:{port}/{database-name}
		url = fmt.Sprintf(
			"user=%s password=%s host=%s port=%s dbname=%s",
			os.Getenv("MYSQL_USER"),
			os.Getenv("MYSQL_PASSWORD"),
			os.Getenv("MYSQL_HOST"),
			os.Getenv("MYSQL_PORT"),
			os.Getenv("MYSQL_NAME"),
		)
	case "redis":
		// URL for Redis connection.
		url = fmt.Sprintf(
			"%s:%s",
			os.Getenv("REDIS_HOST"),
			os.Getenv("REDIS_PORT"),
		)
	case "fiber":
		// URL for Fiber connection.
		url = fmt.Sprintf(
			"%s:%s",
			os.Getenv("SERVER_HOST"),
			os.Getenv("SERVER_PORT"),
		)
	default:
		// Return error message.
		return "", fmt.Errorf("connection name '%v' is not supported", n)
	}

	// Return connection URL.
	return url, nil
}
