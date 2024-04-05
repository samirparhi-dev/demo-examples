package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// Initialize the database connection
	initDB()
	defer db.Close()

	// Handle HTTP requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Extract the user input (unsafe)
		query := r.URL.Query().Get("query")

		// Query the database (vulnerable to SQL injection)
		rows, err := db.Query("SELECT * FROM users WHERE username='" + query + "'")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		// Print the results
		for rows.Next() {
			var id int
			var username string
			err := rows.Scan(&id, &username)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Fprintf(w, "ID: %d, Username: %s\n", id, username)
		}
	})

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func initDB() {
	var err error
	db, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/database")
	if err != nil {
		log.Fatal(err)
	}
}
