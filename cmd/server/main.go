package main

import (
	"context"
	"github.com/word-of-wisdom/internal/server"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	config := server.ParseConfig()
	file, err := os.Open(config.QuotesFilePath)
	mustStart(err)
	defer file.Close()
	quotesRepo, err := server.NewQuotesRepository(file)
	mustStart(err)
	pow := server.NewPow(config.PowDifficulty, config.ChallengeLength)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	signalCh := make(chan os.Signal, 1)
	defer close(signalCh)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		sig := <-signalCh
		slog.Info("Received signal", "signal", sig)
		cancel()
	}()

	s := server.NewServer(config.ListenPort, config.ConnTimeout, quotesRepo, pow)
	err = s.Run(ctx)
	mustStart(err)
}

func mustStart(err error) {
	if err == nil {
		return
	}

	log.Fatalf("Error starting application: %v", err)
}
