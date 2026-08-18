package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"rag-course/app"
	"rag-course/config"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	if err := app.Run(ctx, config.Load()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
