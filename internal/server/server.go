package server

import (
	"bufio"
	"context"
	"fmt"
	"github.com/word-of-wisdom/internal"
	"log/slog"
	"net"
	"strings"
	"sync"
	"time"
)

type quotesRepo interface {
	GetRandomLine() string
}

type pow interface {
	GenerateChallenge() string
	Difficulty() int
	Verify(powChallenge, message string) bool
}

type Server struct {
	listenPort   string
	connDeadline time.Duration
	quotesRepo   quotesRepo
	pow          pow
}

func NewServer(listenPort string, connDeadline time.Duration, quotesRepo quotesRepo, pow pow) *Server {
	return &Server{
		listenPort:   listenPort,
		connDeadline: connDeadline,
		quotesRepo:   quotesRepo,
		pow:          pow,
	}
}

func (s *Server) Run(ctx context.Context) error {
	var lc net.ListenConfig
	listener, err := lc.Listen(ctx, "tcp", ":"+s.listenPort)
	if err != nil {
		return err
	}
	defer listener.Close()
	go func() {
		<-ctx.Done()
		slog.Info("Gracefully shutting down the server...")
		listener.Close()
	}()

	slog.Info(fmt.Sprintf("Server listening on port %s...", s.listenPort))

	var wg sync.WaitGroup

	for {
		conn, err := listener.Accept()
		if err != nil {
			if ctx.Err() != nil {
				slog.Info("Waiting for requests being served...")
				wg.Wait()
				break
			}
			slog.Error("Error accepting connection", "error", err)
			continue
		}

		wg.Add(1)
		go s.handleConnection(ctx, conn, &wg)
	}

	return nil
}

func (s *Server) handleConnection(ctx context.Context, conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			slog.Error("Error closing connection", "error", err)
		}
	}(conn)

	_, err := conn.Write([]byte("You are welcomed on the yet another test task server\n"))
	if err != nil {
		slog.Error("Error writing to connection", "error", err)
		return
	}

	reader := bufio.NewReader(conn)

	err = conn.SetDeadline(time.Now().Add(s.connDeadline))
	if err != nil {
		slog.Error("Error setting connection deadline", "error", err)
		return
	}

	powChallenge := s.pow.GenerateChallenge()
	_, err = conn.Write([]byte(fmt.Sprintf(
		internal.ChallengeString,
		powChallenge,
		s.pow.Difficulty(),
	)))
	if err != nil {
		slog.Error("Error writing to connection", "error", err)
		return
	}

	message, err := reader.ReadString('\n')
	if err != nil {
		if ctx.Err() != nil {
			return
		}
		slog.Error("Error reading string", "error", err, "message", message)
		return
	}
	message = strings.TrimSpace(message)
	if message == "" {
		return
	}
	var response string
	if s.pow.Verify(powChallenge, message) {
		response = s.quotesRepo.GetRandomLine()
	} else {
		response = "Wrong answer"
	}

	_, err = conn.Write([]byte(response + "\n"))
	if err != nil {
		slog.Error("Error writing to the connection", "error", err)
		return
	}

	slog.Info("Connection has been served", "response", response)
}
