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
	ctx, cancel := context.WithTimeout(context.Background(), 300 * time.Millisecond)
	defer cancel()

	reqToServer, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)

	res, err := http.DefaultClient.Do(reqToServer)
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

	bid := string(resCotacaoBody)
	fmt.Println("BID", bid)
}