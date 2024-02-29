package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	reqToServer, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error ao preparar requisição para o server localhost:8080 %s\n", err)
		return
	}

	res, err := http.DefaultClient.Do(reqToServer)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error ao executar chamada GET para /cotacao %s\n", err)
		return
	}
	defer res.Body.Close()

	resCotacaoBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error ao ler Body da request /cotacao %s\n", err)
		return
	}

	bid := string(resCotacaoBody)

	// f, err := os.Create("cotacao.txt")
	f, err := os.OpenFile("cotacao.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error criar arquivo cotacao.txt %s\n", err)
		return
	}

	line := "Dolar: " + bid + "\n"

	_, err = f.WriteString(line)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error adicionar cotação no arquivo cotacao.txt %s\n", err)
		return
	}

	fmt.Println("BID", bid, "Status Code", res.StatusCode)
}
