package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/internal/config"
	"github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/internal/server"
	"google.golang.org/grpc"
)

func main() {
	flagPort := flag.Int("port", 0, "gRPC port (override config)")
	flagReflection := flag.Bool("reflection", false, "enable reflection")
	flagConfig := flag.String("config", "", "config YAML file")
	flag.Parse()

	if *flagConfig != "" {
		os.Setenv("CONFIG_FILE", *flagConfig)
	}

	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}
	if *flagPort != 0 {
		conf.Server.Port = *flagPort
	}
	if *flagReflection {
		conf.Server.EnableReflection = true
	}

	srv, err := server.New(conf)
	if err != nil {
		log.Fatalf("❌ Failed to initialize server: %v", err)
	}

	log.Printf("🚀 Serving gRPC on %s (reflection=%v)", srv.Addr(), conf.Server.EnableReflection)

	errCh := make(chan error, 1)
	go func() {
		if err := srv.Serve(); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			errCh <- err
		}
	}()

	// Handle SIGINT/SIGTERM for graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	select {
	case sig := <-sigCh:
		log.Printf("📴 Caught %s – gracefully shutting down…", sig)
		srv.GracefulStop()
		log.Println("✅ Server stopped cleanly.")
	case err := <-errCh:
		log.Fatalf("❌ Server error: %v", err)
	}
}
