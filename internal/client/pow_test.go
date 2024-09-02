package client

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/word-of-wisdom/internal/server"
	"testing"
	"time"
)

func TestFindSolution_Success(t *testing.T) {
	pow := NewPow(1)
	challenge := "testchallenge"
	difficulty := 1

	solution, err := pow.FindSolution(context.Background(), challenge, difficulty)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	assert.Equal(t, solution, "300")
}

func TestFindSolution_Timeout(t *testing.T) {
	pow := NewPow(1)
	challenge := "testchallenge"
	difficulty := 10

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	_, err := pow.FindSolution(ctx, challenge, difficulty)
	if err == nil {
		t.Fatalf("Expected an error due to timeout, got nil")
	}

	if err.Error() != "timeout finding solution" {
		t.Fatalf("Expected timeout error, got %v", err)
	}
}

func TestFindSolution_MultipleGoroutines(t *testing.T) {
	concurrency := 4
	pow := NewPow(concurrency)
	challenge := "testchallenge"
	difficulty := 2
	serverPow := server.NewPow(difficulty, 0)

	solution, err := pow.FindSolution(context.Background(), challenge, difficulty)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	
	assert.True(t, serverPow.Verify(challenge, solution))
}
