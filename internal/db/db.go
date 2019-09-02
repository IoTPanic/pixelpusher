package db

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// Connect starts the SQLite connection
func Connect(dbPath string) error {
	var err error // This is to satisfy VSCode's highlighting
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Println("Failed to connect to the database! - ", err.Error())
		return err
	}
	log.Println("Connected to DB")
	return CreateTables()
}

func CreateTables() error {
	statement, err := db.Prepare(createFixturesTableStatement)
	if err != nil {
		log.Println("SQLite error", err.Error())
		return err
	}
	statement.Exec()
	statement, err = db.Prepare(createChannelsTableStatement)
	if err != nil {
		log.Println("SQLite error", err.Error())
		return err
	}
	statement.Exec()
	return nil
}

func Close() {
	db.Close()
}

func Test() {
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	if err != nil {
		log.Println(err.Error())
	}
	statement.Exec()
	statement, _ = db.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
	statement.Exec("Nic", "Raboy")
	rows, _ := db.Query("SELECT id, firstname, lastname FROM people")
	var id int
	var firstname string
	var lastname string
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	}
}
