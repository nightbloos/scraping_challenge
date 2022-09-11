package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"scraping_challenge/services/cometco-scraper/app"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	if err := app.NewApplication().Run(ctx); err != nil {
		log.Fatal(err)
	}
}
