package server

import (
	"reflect"
	"testing"
	"time"
)

func TestParseConfig(t *testing.T) {
	tests := []struct {
		name           string
		expectedConfig *AppConfig
	}{
		{
			name: "default values",
			expectedConfig: &AppConfig{
				ListenPort:      "8080",
				ConnTimeout:     10 * time.Second,
				PowDifficulty:   6,
				ChallengeLength: 20,
				QuotesFilePath:  "./internal/server/quotes_dump.txt",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseConfig()
			if !reflect.DeepEqual(got, tt.expectedConfig) {
				t.Errorf("ParseConfig() = %v, want %v", got, tt.expectedConfig)
			}
		})
	}
}
