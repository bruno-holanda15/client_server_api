package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"server_api_go_expert/infra/repositories"
	"time"
)

type CotacaoDolarHTTP struct {
	db *sql.DB
}

type DolarInfo struct {
	USDBRL struct {
		Bid string `json:"bid"`
	} `json:"USDBRL"`
}

func NewCotacaoDolarHTTP(db *sql.DB) *CotacaoDolarHTTP {
	return &CotacaoDolarHTTP{
		db: db,
	}
}	

func (c *CotacaoDolarHTTP) CotacaoDolar(res http.ResponseWriter, req *http.Request) {
	ctxReqDolar, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	reqDolar, err := http.NewRequestWithContext(ctxReqDolar, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error ao criar request com context %s\n", err)
		return
	}

	respDolar, err := http.DefaultClient.Do(reqDolar)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error ao executar chamada GET para economia.awesomeapi %s\n", err)
		return
	}

	defer respDolar.Body.Close()
	bodyDolar, err := io.ReadAll(respDolar.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error ao recuperar Body %s\n", err)
		return
	}

	var dolarInfo DolarInfo
	err = json.Unmarshal(bodyDolar, &dolarInfo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error unmarshal para struct DolarInfo %s\n", err)
		return
	}

	ctxDB, cancelDB := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancelDB()

	repository := repositories.NewServerRepository(c.db)

	err = repository.Insert(ctxDB, "dolar", dolarInfo.USDBRL.Bid)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error ao persistir bid no banco de dados %s\n", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	res.Write([]byte(dolarInfo.USDBRL.Bid))
}
