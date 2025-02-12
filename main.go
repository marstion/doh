package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/marstion/doh_client/internal/cli"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
	)
	defer cancel()

	err := cli.CommandRoot.ExecuteContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
