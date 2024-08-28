package server

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
)

type QuotesRepo struct {
	quotes []string
}

func NewQuotesRepository(reader io.Reader) (*QuotesRepo, error) {
	repo := &QuotesRepo{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		repo.quotes = append(repo.quotes, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(repo.quotes) == 0 {
		return nil, fmt.Errorf("file is empty")
	}

	return repo, nil
}

func (r *QuotesRepo) GetRandomLine() string {
	return r.quotes[rand.Intn(len(r.quotes))] + "\n"
}
