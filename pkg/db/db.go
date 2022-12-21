package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func init() {
	var err error
	_, err = os.Stat("./storage.db")
	if err != nil {
		DB, err = sql.Open("sqlite3", "./storage.db")
		if err != nil {
			log.Fatal(err)
		}
		sqlStmt := `
	create table files (id INTEGER PRIMARY KEY AUTOINCREMENT, filepath text, hash text, CONSTRAINT filehash UNIQUE (filepath, hash));
	`
		_, err = DB.Exec(sqlStmt)
		if err != nil {
			log.Fatalf("%q: %s\n", err, sqlStmt)
		}
	} else {
		DB, err = sql.Open("sqlite3", "./storage.db")
		if err != nil {
			log.Fatal(err)
		}
	}

}

func SaveFileHash(filePath string, fileHash string) error {
	tx, err := DB.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into files(filepath, hash) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(filePath, fileHash)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
