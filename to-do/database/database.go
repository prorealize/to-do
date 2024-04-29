package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func ConnectDB() {
	var err error
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	for _, env := range []string{host, port, user, password, dbname} {
		if env == "" {
			log.Fatal("Not all environment variable are set.")
		}
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	Db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	err = Db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Ensure the table exists
	createItemsTable(err, Db)
	log.Printf("Database successfully connected!")
}

func createItemsTable(err error, Db *sql.DB) {
	createTable := `CREATE TABLE IF NOT EXISTS items (
    		id SERIAL PRIMARY KEY,
    		title TEXT,
    		description TEXT,
    		status TEXT DEFAULT 'new'
	);`
	_, err = Db.Exec(createTable)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
}
