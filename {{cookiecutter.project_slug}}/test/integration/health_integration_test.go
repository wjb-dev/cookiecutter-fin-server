package integration

import (
	"context"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"testing"
	"time"
)

func TestHealthCheckIntegration(t *testing.T) {
	conn, cleanup, err := dialer()
	if err != nil {
		t.Fatalf("dial bufconn: %v", err)
	}
	defer cleanup()

	healthClient := healthpb.NewHealthClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := healthClient.Check(ctx, &healthpb.HealthCheckRequest{
		Service: "v1.EchoService",
	})
	if err != nil {
		t.Fatalf("HealthCheck RPC failed: %v", err)
	}
	if resp.Status != healthpb.HealthCheckResponse_SERVING {
		t.Errorf("Health status = %v; want SERVING", resp.Status)
	}
}
