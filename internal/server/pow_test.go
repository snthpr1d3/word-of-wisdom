package server

import (
	"testing"
)

func TestNewPow(t *testing.T) {
	tests := []struct {
		name                    string
		customDifficulty        int
		customChallengeLength   int
		expectedDifficulty      int
		expectedChallengeLength int
	}{
		{
			name:                    "default values",
			customDifficulty:        0,
			customChallengeLength:   0,
			expectedDifficulty:      defaultDifficulty,
			expectedChallengeLength: defaultChallengeLength,
		},
		{
			name:                    "custom values within range",
			customDifficulty:        5,
			customChallengeLength:   30,
			expectedDifficulty:      5,
			expectedChallengeLength: 30,
		},
		{
			name:                    "custom values out of range",
			customDifficulty:        100, // out of range
			customChallengeLength:   200, // out of range
			expectedDifficulty:      defaultDifficulty,
			expectedChallengeLength: defaultChallengeLength,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pow := NewPow(tt.customDifficulty, tt.customChallengeLength)
			if pow.Difficulty() != tt.expectedDifficulty {
				t.Errorf("NewPow() difficulty = %v, want %v", pow.Difficulty(), tt.expectedDifficulty)
			}
			if pow.challengeLength != tt.expectedChallengeLength {
				t.Errorf("NewPow() challengeLength = %v, want %v", pow.challengeLength, tt.expectedChallengeLength)
			}
		})
	}
}

func TestGenerateChallenge(t *testing.T) {
	pow := NewPow(4, 20)

	challenge := pow.GenerateChallenge()

	if len(challenge) != 20 {
		t.Errorf("GenerateChallenge() length = %v, want %v", len(challenge), 20)
	}
}

func TestVerify(t *testing.T) {
	pow := NewPow(6, 20)

	if !pow.Verify("Hkyip2oDKcC7dCaR~%cD", "354baf18") {
		t.Errorf("Verify() failed for challenge %v with solution %v", "Hkyip2oDKcC7dCaR~%cD", "354baf18")
	}

	if pow.Verify("whatever", "invalidsolution") {
		t.Errorf("Verify() should have failed for invalid solution")
	}
}
