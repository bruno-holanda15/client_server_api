package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/cotacao", CotacaoDolar)
	http.ListenAndServe(":8080", nil)

}

func CotacaoDolar(res http.ResponseWriter, req *http.Request) {
	respDolar, err := http.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error ao executar chamada GET para economia.awesomeapi %s", err)
		return
	}

	defer respDolar.Body.Close()
	bodyDolar, err := io.ReadAll(respDolar.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error ao recuperar Body %s", err)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	res.Write(bodyDolar)
}
