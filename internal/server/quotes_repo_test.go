package server

import (
	"strings"
	"testing"
)

func TestNewQuotesRepository(t *testing.T) {
	quotes := "Quote 1\nQuote 2\nQuote 3\n"
	reader := strings.NewReader(quotes)

	repo, err := NewQuotesRepository(reader)
	if err != nil {
		t.Fatalf("NewQuotesRepository() error = %v", err)
	}

	expectedQuotes := []string{"Quote 1", "Quote 2", "Quote 3"}
	if len(repo.quotes) != len(expectedQuotes) {
		t.Errorf("NewQuotesRepository() got %d quotes, want %d", len(repo.quotes), len(expectedQuotes))
	}
	for i, quote := range expectedQuotes {
		if repo.quotes[i] != quote {
			t.Errorf("NewQuotesRepository() quote %d = %v, want %v", i, repo.quotes[i], quote)
		}
	}
}

func TestGetRandomLine(t *testing.T) {
	quotes := "Quote 1\nQuote 2\nQuote 3\n"
	reader := strings.NewReader(quotes)

	repo, err := NewQuotesRepository(reader)
	if err != nil {
		t.Fatalf("NewQuotesRepository() error = %v", err)
	}

	lineCount := map[string]int{}
	for i := 0; i < 100; i++ {
		line := repo.GetRandomLine()
		lineCount[line]++
	}

	expectedQuotes := []string{"Quote 1\n", "Quote 2\n", "Quote 3\n"}
	for _, quote := range expectedQuotes {
		if lineCount[quote] == 0 {
			t.Errorf("GetRandomLine() did not return quote %v", quote)
		}
	}
}

func TestNewQuotesRepository_EmptyReader(t *testing.T) {
	reader := strings.NewReader("")

	_, err := NewQuotesRepository(reader)
	if err == nil {
		t.Fatal("NewQuotesRepository() should return an error for empty reader")
	}
}
