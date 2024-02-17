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
	res, err := http.Get("http://localhost:8080/cotacao")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error ao executar chamada GET para /cotacao %s", err)
		return
	}
	defer res.Body.Close()

	resCotacaoBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error ao ler Body da request /cotacao %s", err)
		return
	}

	var dolarInfo DolarInfo

	err = json.Unmarshal(resCotacaoBody, &dolarInfo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error ao transformar body para DolarInfo struct %s", err)
		return
	}

	fmt.Println("Executado com sucesso", dolarInfo.USDBRL.Bid)
}