package server

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strings"
)

const (
	defaultDifficulty      = 3
	maxDifficulty          = 31
	defaultChallengeLength = 20
	maxChallengeLength     = 100
	asciiChars             = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:',.<>?/`~"
)

type Pow struct {
	difficulty      int
	challengeLength int
}

func NewPow(customDifficulty, customChallengeLength int) *Pow {
	difficulty := defaultDifficulty
	if customDifficulty > 1 && customDifficulty <= maxDifficulty {
		difficulty = customDifficulty
	}
	challengeLength := defaultChallengeLength
	if customChallengeLength > 0 && customChallengeLength <= maxChallengeLength {
		challengeLength = customChallengeLength
	}
	return &Pow{difficulty: difficulty, challengeLength: challengeLength}
}

func (p *Pow) GenerateChallenge() string {
	result := make([]byte, p.challengeLength)
	asciiBytes := []byte(asciiChars)

	for i := range result {
		result[i] = asciiBytes[rand.Intn(len(asciiBytes))]
	}

	return string(result)
}

func (p *Pow) Verify(challenge string, solution string) bool {
	result := challenge + solution
	hash := sha256.Sum256([]byte(result))
	hashStr := hex.EncodeToString(hash[:])
	return strings.HasPrefix(hashStr, strings.Repeat("0", p.difficulty))
}

func (p *Pow) Difficulty() int {
	return p.difficulty
}
