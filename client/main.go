package main

import (
	"client_api_go_expert/app"
	"context"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	bid, err := app.GetBid(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = app.WriteResponseIntoFile(bid)
	if err != nil {
		log.Fatal(err)
		return
	}
}
