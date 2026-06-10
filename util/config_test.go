package util

import "testing"

func TestLoadConfigFromEnvironmentWithoutConfigFile(t *testing.T) {
	t.Setenv("DB_SOURCE", "test-db-source")
	t.Setenv("SERVER_ADDRESS", "127.0.0.1:9999")

	config, err := LoadConfig(t.TempDir())
	if err != nil {
		t.Fatalf("LoadConfig returned error: %v", err)
	}

	if config.DBSource != "test-db-source" {
		t.Fatalf("expected DB_SOURCE from environment, got %q", config.DBSource)
	}

	if config.ServerAddress != "127.0.0.1:9999" {
		t.Fatalf("expected SERVER_ADDRESS from environment, got %q", config.ServerAddress)
	}
}
