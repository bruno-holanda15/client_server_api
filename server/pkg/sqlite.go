package pkg

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)


func ConnectionSQLite() *sql.DB{
	db, err := sql.Open("sqlite3", "../server_database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return db
}
