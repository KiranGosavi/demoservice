package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name          string
		serviceConfig string
		wantErr       bool
	}{
		{
			name:          "wrong_port",
			serviceConfig: `{"demo_service_port": "8080"", "log_level": "info"}`,
			wantErr:       true,
		},
		{
			name:          "wrong_logLevel",
			serviceConfig: `{"demo_service_port": 8080, "log_level": ["wrong"]}`,
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		err := os.Setenv("DEMO_SERVER_PROPERTIES", tt.serviceConfig)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		t.Run(tt.name, func(t *testing.T) {
			if err = Load(); (err != nil) != tt.wantErr {
				t.Errorf("Load() error: %v, want error: %v", err, tt.wantErr)
			}

		})
	}
}
