package pkg

import (
	"database/sql"
	"os"
	_ "github.com/mattn/go-sqlite3"
)

func ConnectionSQLite() *sql.DB{
	os.Remove("./server_database.db")

	db, err := sql.Open("sqlite3", "./server_database.db")
	if err != nil {
		panic(err)
	}

	return db
}
