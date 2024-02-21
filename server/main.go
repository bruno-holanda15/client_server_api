package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {

	http.HandleFunc("/cotacao", CotacaoDolar)
	http.ListenAndServe(":8080", nil)

}

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

func CotacaoDolar(res http.ResponseWriter, req *http.Request) {
	ctxReqDolar, cancel := context.WithTimeout(context.Background(), 200 * time.Millisecond)
	defer cancel()

	reqDolar, err := http.NewRequestWithContext(ctxReqDolar, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error ao criar request com context %s", err)
		return
	}

	respDolar, err := http.DefaultClient.Do(reqDolar)
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
		fmt.Fprintf(os.Stderr, "Error unmarshal para struct DolarInfo %s", err)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	res.Write([]byte(dolarInfo.USDBRL.Bid))
}
