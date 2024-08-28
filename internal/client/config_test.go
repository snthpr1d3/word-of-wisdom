package client

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
				Concurrency:    0,
				ServerAddr:     "localhost:8080",
				ConnTimeout:    10 * time.Second,
				SolvingTimeout: 5 * time.Second,
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
