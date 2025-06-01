package db

import (
	"database/sql"
	"log"

	_ "github.com/glebarez/sqlite"
)

// public static Connection initDB() in Java
func InitDB() *sql.DB {

	// returns multiple values, the result in db and if there is an error in err
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal("Failed to open DB:", err)
	}

// ":=" is an operator for assigning 
	createTableQuery := `
	CREATE TABLE documents (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT
	);`

	// when i don't care about the result I assign it to "_"
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}

	return db
}
