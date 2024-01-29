package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

//
// 	NOTE: panic functions are used because there is no point in running the program if we cannot
//  			store the stock information
//

// Global database variable
var DB *sql.DB

func InitDB() {

	var err error

	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Database connection failed.")
	}

	// 5 is an arbitrary number here, subject to change
	DB.SetMaxOpenConns(5)
	DB.SetMaxIdleConns(5)

	createTables()

}

func createTables() {
	createStockTable := `
	CREATE TABLE IF NOT EXISTS stocks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		ticker TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		value REAL NOT NULL
	)
	`

	_, err := DB.Exec(createStockTable)

	if err != nil {
		fmt.Println(err)
		panic("Could not create stock table.")
	}

}
