package server_test

import (
	"context"
	"testing"

	"github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/internal/server"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func TestNewHealthServer_Statuses(t *testing.T) {
	hs := server.NewHealthServer()

	tests := []struct {
		name    string
		service string
		want    grpc_health_v1.HealthCheckResponse_ServingStatus
	}{
		{"overall", "", grpc_health_v1.HealthCheckResponse_SERVING},
		{"EchoService", "v1.EchoService", grpc_health_v1.HealthCheckResponse_SERVING},
		{"PingService", "v1.PingService", grpc_health_v1.HealthCheckResponse_SERVING},
	}

	for _, tc := range tests {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			resp, err := hs.Check(context.Background(), &grpc_health_v1.HealthCheckRequest{
				Service: tc.service,
			})
			if err != nil {
				t.Fatalf("Check(%q) returned error: %v", tc.service, err)
			}
			if resp.Status != tc.want {
				t.Errorf("Check(%q) = %v; want %v", tc.service, resp.Status, tc.want)
			}
		})
	}
}
