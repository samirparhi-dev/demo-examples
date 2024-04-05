package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    var userInput string

    // Taking user input
    fmt.Print("Enter your username: ")
    fmt.Scanln(&userInput)

    // Simulating SQL query
    query := fmt.Sprintf("SELECT * FROM users WHERE username='%s'", userInput)

    // Executing SQL query (This is vulnerable to SQL injection)
    db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/database")
    if err != nil {
 	   panic(err.Error())
    }
    defer db.Close()

    rows, err := db.Query(query)
    if err != nil {
 	   panic(err.Error())
    }
    defer rows.Close()

    // Displaying results
    for rows.Next() {
 	   var id int
 	   var username string
 	   rows.Scan(&id, &username)
 	   fmt.Printf("ID: %d, Username: %s\n", id, username)
    }
}
