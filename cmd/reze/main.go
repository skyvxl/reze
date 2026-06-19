package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/skyvxl/reze/internal/cli"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	rootCmd := cli.NewRootCommand()

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return
		}
		stop()
		os.Exit(1)
	}
	stop()
}
