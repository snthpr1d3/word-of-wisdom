package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/word-of-wisdom/internal"
	"github.com/word-of-wisdom/internal/client"
	"log/slog"
	"net"
	"time"
)

func main() {
	ctx := context.Background()
	config := client.ParseConfig()
	conn, err := connectToServer(config.ServerAddr, config.ConnTimeout)
	if err != nil {
		slog.Error("Error connecting to server", "error", err)
		return
	}
	defer conn.Close()

	start := time.Now()
	solution, err := performChallenge(ctx, conn, config.SolvingTimeout)
	if err != nil {
		slog.Error("Error finding solution", "error", err)
		return
	}
	duration := time.Since(start)
	slog.Info("Solution found", "solution", solution, "duration", duration)

	err = handleServerResponse(conn, solution)
	if err != nil {
		slog.Error("Error handling server response", "error", err)
		return
	}
}

func connectToServer(serverAddr string, connTimeout time.Duration) (net.Conn, error) {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		return nil, err
	}

	err = conn.SetDeadline(time.Now().Add(connTimeout))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func performChallenge(ctx context.Context, conn net.Conn, solvingTimeout time.Duration) (string, error) {
	reader := bufio.NewReader(conn)

	welcomeMsg, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	slog.Info(welcomeMsg)

	challengeMsg, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	slog.Info(challengeMsg)

	var challenge string
	var difficulty int
	_, err = fmt.Sscanf(challengeMsg, internal.ChallengeString, &challenge, &difficulty)
	if err != nil {
		return "", err
	}

	pow := client.NewPow(0)
	ctxWithTimeout, cancel := context.WithTimeout(ctx, solvingTimeout)
	defer cancel()

	return pow.FindSolution(ctxWithTimeout, challenge, difficulty)
}

func handleServerResponse(conn net.Conn, solution string) error {
	_, err := conn.Write([]byte(solution + "\n"))
	if err != nil {
		return err
	}

	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return err
	}

	slog.Info(response)
	return nil
}
