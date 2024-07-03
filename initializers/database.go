package initializers

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/denisenkom/go-mssqldb"
)

func InitDB() {

	database := os.Getenv("DB_DATABASE")
	host := os.Getenv("DB_HOST")
	portStr := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USER")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Invalid port number: %v", err)
	}

	// password := url.QueryEscape(fmt.Sprintf("%s#$%s@#$%d", os.Getenv("DBPassword1"), os.Getenv("DBPassword2"), os.Getenv("DBPassword3")))
	password := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&encrypt=true&TrustServerCertificate=true", username, password, host, port, database)

	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		log.Fatalf("Failed to open a DB connection: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}

}
