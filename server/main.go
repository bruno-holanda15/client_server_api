package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"server_api_go_expert/infra/handlers"
	"server_api_go_expert/infra/repositories"
	"server_api_go_expert/pkg"
)

var db *sql.DB

func init() {
	setupDatabase()
}

func main() {

	cotacaoDolar := handlers.NewCotacaoDolarHTTP(db)
	defer db.Close()

	http.HandleFunc("/cotacao", cotacaoDolar.CotacaoDolar)
	http.ListenAndServe(":8080", nil)

}

func setupDatabase() {
	db = pkg.ConnectionSQLite()

	err := repositories.CreateTable(db)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error criando table cotacoes %s\n", err)
		panic(err)
	}
}