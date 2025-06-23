package server

import (
	"fmt"
	"net"

	"github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/internal/config"
	"github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/internal/handler"
	v1 "github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/pkg/proto/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

// Server wraps a gRPC server and its listener.
type Server struct {
	*grpc.Server
	lis net.Listener
}

// New constructs a Server with all services registered and ready to serve.
func New(conf *config.Config) (*Server, error) {
	host := conf.Server.Host
	if host == "" {
		host = "0.0.0.0"
	}
	port := conf.Server.Port
	if port == 0 {
		port = 50051
	}
	addr := fmt.Sprintf("%s:%d", host, port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("listen %s: %w", addr, err)
	}

	var opts []grpc.ServerOption

	if conf.TLS.Enabled {
		creds, err := credentials.NewServerTLSFromFile(
			conf.TLS.CertFile, conf.TLS.KeyFile,
		)
		if err != nil {
			return nil, fmt.Errorf("load TLS: %w", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}
	if conf.Server.MaxRecvMsgSize > 0 {
		opts = append(opts, grpc.MaxRecvMsgSize(conf.Server.MaxRecvMsgSize))
	}
	if conf.Server.MaxSendMsgSize > 0 {
		opts = append(opts, grpc.MaxSendMsgSize(conf.Server.MaxSendMsgSize))
	}

	s := grpc.NewServer(opts...)

	// Register gRPC services
	v1.RegisterEchoServiceServer(s, handler.NewEchoServiceServer())
	v1.RegisterPingServiceServer(s, handler.NewPingServiceServer())

	// Health check
	healthpb.RegisterHealthServer(s, NewHealthServer())

	// Reflection
	if conf.Server.EnableReflection {
		reflection.Register(s)
	}

	return &Server{Server: s, lis: lis}, nil
}

// Serve starts serving (blocking).
func (s *Server) Serve() error {
	return s.Server.Serve(s.lis)
}

// Addr returns the listener address (useful for logging).
func (s *Server) Addr() string {
	return s.lis.Addr().String()
}
