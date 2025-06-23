package config

import (
	"github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/internal/config"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfigFromFile_WithEnvOverrides(t *testing.T) {
	// --- arrange ----------------------------------------------------------------
	tmpDir := t.TempDir()
	cfgPath := filepath.Join(tmpDir, "test.yaml")

	yamlContent := `
server:
  host: 127.0.0.1
  port: 12345
  enable_reflection: true
tls:
  enabled: false
logging:
  level: info
metrics:
  enabled: true
  port: 9100
`
	if err := os.WriteFile(cfgPath, []byte(yamlContent), 0o600); err != nil {
		t.Fatalf("failed to write temp config: %v", err)
	}

	t.Setenv("GRPC_PORT", "54321")
	t.Setenv("TLS_ENABLED", "true")
	t.Setenv("LOG_LEVEL", "debug")

	// --- act --------------------------------------------------------------------
	got, err := config.LoadConfigFromFile(cfgPath) // direct call to unexported func
	if err != nil {
		t.Fatalf("loadConfigFromFile(%s) error = %v", cfgPath, err)
	}

	// --- assert -----------------------------------------------------------------
	if got.Server.Port != 54321 {
		t.Errorf("Server.Port = %d; want 54321", got.Server.Port)
	}
	if !got.TLS.Enabled {
		t.Error("TLS.Enabled = false; want true")
	}
	if got.Logging.Level != "debug" {
		t.Errorf("Logging.Level = %q; want %q", got.Logging.Level, "debug")
	}
	if got.Metrics.Port != 9100 {
		t.Errorf("Metrics.Port = %d; want 9100", got.Metrics.Port)
	}
}
