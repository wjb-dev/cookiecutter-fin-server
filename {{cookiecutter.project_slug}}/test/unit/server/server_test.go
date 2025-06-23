package server_test

import (
	"net"
	"strings"
	"testing"

	"github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/internal/config"
	"github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/internal/server"
)

func TestNewServer_BindsToExpectedAddr(t *testing.T) {
	conf := &config.Config{
		Server: config.ServerConfig{
			Host:             "127.0.0.1",
			Port:             50551,
			EnableReflection: false,
		},
		TLS: config.TLSConfig{
			Enabled: false,
		},
	}

	srv, err := server.New(conf)
	if err != nil {
		t.Fatalf("New() returned error: %v", err)
	}
	defer srv.GracefulStop()

	addr := srv.Addr()
	if !strings.Contains(addr, "127.0.0.1:50551") {
		t.Errorf("Expected address to contain '127.0.0.1:50551', got %q", addr)
	}

	// Ensure it's a valid net.Addr
	if _, err := net.ResolveTCPAddr("tcp", addr); err != nil {
		t.Errorf("Addr() returned invalid TCP address: %v", err)
	}
}
