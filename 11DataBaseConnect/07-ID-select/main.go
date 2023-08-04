package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Establish a connection
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/storeDB")
	if err != nil {
		fmt.Println("Cannot connect to the database:", err)
		return
	}
	defer db.Close()

	// Query the database
	var num int
	query := "SELECT id FROM table1 WHERE email = ?"  // remember you must have id in database
	err = db.QueryRow(query, "monish@gmail.com").Scan(&num)
	if err != nil {
		fmt.Println("Error while searching for email:", err)
		return
	}

	fmt.Println("User ID:", num)
}
