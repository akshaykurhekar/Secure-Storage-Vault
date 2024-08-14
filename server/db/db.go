package db

import (
	"database/sql"
	"log"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Init(){

	DB, err := sql.Open("sqlite","wallet.db")
	if err != nil {
		log.Fatal("Failed to connect DB...!")
	}
	defer DB.Close()
	// create vault table - query
	credentialsTable := `CREATE TABLE IF NOT EXISTS credentials (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	vid INTEGER,
	credential TEXT);`

	// Execute the SQL statement
    statement, err := DB.Prepare(credentialsTable)
    if err != nil {
        panic(err)
    }
    defer statement.Close()

    _, err = statement.Exec()
    if err != nil {
        panic(err)
    }


	// Exec vault table
	if _, err := DB.Exec(credentialsTable); err != nil {
		log.Fatal("Vault Table Exec Failed..")
	}

	// create credentials table
	vaultTable := `CREATE TABLE IF NOT EXISTS vault (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		password TEXT,
		desc TEXT);`
	
		// exec credential table
	if _, err := DB.Exec(vaultTable); err != nil {
		log.Fatal("Vault Table Exec Failed..")
	}

	log.Println("Database setup success..!")
}