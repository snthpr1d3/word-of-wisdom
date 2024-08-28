package server

import (
	"flag"
	"time"
)

type AppConfig struct {
	ListenPort      string
	ConnTimeout     time.Duration
	PowDifficulty   int
	ChallengeLength int
	QuotesFilePath  string
}

func ParseConfig() *AppConfig {
	listenPort := flag.String("port", "8080", "the listen port number")
	connTimeout := flag.Duration("conn-timeout", 10*time.Second, "connection timeout")
	powDifficulty := flag.Int("pow-difficulty", 5, "difficulty of pow operation")
	challengeLength := flag.Int("challenge-length", 20, "the length of a challenge word")
	quotesFilePath := flag.String(
		"quotes-file-path", "./internal/server/quotes_dump.txt", "path to the file with quotes")

	flag.Parse()

	return &AppConfig{
		ListenPort:      *listenPort,
		ConnTimeout:     *connTimeout,
		PowDifficulty:   *powDifficulty,
		ChallengeLength: *challengeLength,
		QuotesFilePath:  *quotesFilePath,
	}
}
