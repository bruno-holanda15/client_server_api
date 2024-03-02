package main

import (
	"net/http"
	"server_api_go_expert/infra/handlers"
)

func main() {

	cotacaoDolar := handlers.CotacaoDolarHTTP{}

	http.HandleFunc("/cotacao", cotacaoDolar.CotacaoDolar)
	http.ListenAndServe(":8080", nil)

}