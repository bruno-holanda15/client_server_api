package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type DolarInfo struct {
	USDBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

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

	var dolarInfo DolarInfo

	err = json.Unmarshal(bodyDolar, &dolarInfo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error ao transformar body para DolarInfo struct %s", err)
			return
	}

	fmt.Println(dolarInfo)
}
