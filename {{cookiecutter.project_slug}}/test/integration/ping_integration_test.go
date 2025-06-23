package integration

import (
	"context"
	v1 "github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/pkg/proto/v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"testing"
	"time"
)

func TestPingIntegration(t *testing.T) {
	conn, cleanup, err := dialer() // reuse dialer() from Echo test
	if err != nil {
		t.Fatalf("dial bufconn: %v", err)
	}
	defer cleanup()

	client := v1.NewPingServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.Ping(ctx, &emptypb.Empty{})
	if err != nil {
		t.Fatalf("Ping RPC failed: %v", err)
	}

	if resp.Status != "SERVING" {
		t.Errorf("Status = %q; want %q", resp.Status, "SERVING")
	}
	if resp.Service != "{{ cookiecutter.project_slug }}" {
		t.Errorf("Service = %q; want %q", resp.Service, "{{ cookiecutter.project_slug }}")
	}
}
