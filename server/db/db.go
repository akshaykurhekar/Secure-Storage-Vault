package db

import (
	"database/sql"
	"log"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Init(){
	var err error
	DB, err := sql.Open("sqlite","wallet.db")
	if err != nil {
		log.Fatal("Failed to connect DB...!")
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
		return
	}

	// create vault table - query
	credentialsTable := `CREATE TABLE IF NOT EXISTS credentials (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	vid INTEGER,
	credential TEXT);`

	// Execute the SQL statement
    statement, err := DB.Prepare(credentialsTable)
    if err != nil {
        log.Fatal("Error in Statement prepare",err)
    }
   
    _, err = statement.Exec()
    if err != nil {
        log.Fatal(err)
    }

	// Exec vault table
	if _, err := DB.Exec(credentialsTable); err != nil {
		log.Fatal("Credential Table Exec Failed..")
		return
	}

	log.Println("Database setup success..!")
}

func OpenConn()(*sql.DB) {
	DB, err := sql.Open("sqlite","wallet.db")
	if err != nil {
		log.Fatal("Failed to connect DB...!")
	}
	return DB
}