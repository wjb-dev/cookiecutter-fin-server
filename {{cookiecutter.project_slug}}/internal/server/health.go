package server

import (
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func NewHealthServer() *health.Server {
	hs := health.NewServer()

	// Overall server process
	hs.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)

	// Fine-grained: specific gRPC services you registered
	hs.SetServingStatus("v1.EchoService", healthpb.HealthCheckResponse_SERVING)

	hs.SetServingStatus("v1.PingService", healthpb.HealthCheckResponse_SERVING)

	return hs
}
