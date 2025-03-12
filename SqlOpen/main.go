package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "tasks.db")
	if err != nil {
		fmt.Println("Error opening database", err)
	}
	db.Ping()
	defer db.Close()
}
