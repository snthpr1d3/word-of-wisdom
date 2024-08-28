package client

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"log/slog"
	"math"
	"runtime"
	"strconv"
	"strings"
)

type Pow struct {
	concurrency int
}

func NewPow(customConcurrency int) *Pow {
	concurrency := runtime.NumCPU()
	if customConcurrency > 0 {
		concurrency = customConcurrency
	}
	return &Pow{
		concurrency: concurrency,
	}
}

func (p *Pow) FindSolution(ctx context.Context, challenge string, difficulty int) (string, error) {
	slog.Info("Concurrency has been set as", "concurrency", p.concurrency)

	prefix := strings.Repeat("0", difficulty)

	result := make(chan string)
	defer close(result)
	ctxWithCancel, cancel := context.WithCancel(ctx)
	defer cancel()

	for i := 0; i < p.concurrency; i++ {
		go func(ctx context.Context, threadNum int, result chan<- string) {
			threadPrefix := strconv.Itoa(threadNum)
			for i := 0; i < math.MaxInt; i++ {
				select {
				case <-ctx.Done():
					return
				default:
					solution := fmt.Sprintf("%x%x", threadPrefix, i)
					data := challenge + solution
					hash := sha256.Sum256([]byte(data))
					hashStr := hex.EncodeToString(hash[:])
					if strings.HasPrefix(hashStr, prefix) {
						select {
						case result <- solution:
							cancel()
						default:
						}
						return
					}
				}
			}
		}(ctxWithCancel, i, result)
	}

	select {
	case solution, ok := <-result:
		if !ok {
			return "", errors.New("no solution found")
		}
		return solution, nil
	case <-ctx.Done():
		return "", errors.New("timeout finding solution")
	}
}
