package client

import (
	"flag"
	"time"
)

type AppConfig struct {
	Concurrency    int
	ServerAddr     string
	ConnTimeout    time.Duration
	SolvingTimeout time.Duration
}

func ParseConfig() *AppConfig {
	concurrency := flag.Int("concurrency", 0, "concurrency of the solving algorithm")
	serverAddr := flag.String("server-address", "localhost:8080", "the server address")
	connTimeout := flag.Duration("conn-timeout", 10*time.Second, "connection deadline")
	solvingTimeout := flag.Duration("solving-timeout", 5*time.Second, "solving deadline")

	flag.Parse()

	return &AppConfig{
		Concurrency:    *concurrency,
		ServerAddr:     *serverAddr,
		ConnTimeout:    *connTimeout,
		SolvingTimeout: *solvingTimeout,
	}
}
