package integration

import (
	"context"
	"testing"
	"time"

	v1 "github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/pkg/proto/v1"
)

func TestEchoIntegration(t *testing.T) {
	conn, cleanup, err := dialer()
	if err != nil {
		t.Fatalf("failed to dial bufconn: %v", err)
	}
	defer cleanup()

	client := v1.NewEchoServiceClient(conn)

	// Set a deadline to avoid hangs
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Perform the Echo RPC
	req := &v1.EchoRequest{Message: "hello integration"}
	resp, err := client.Echo(ctx, req)
	if err != nil {
		t.Fatalf("Echo RPC failed: %v", err)
	}
	if resp.Message != req.Message {
		t.Errorf("expected message %q, got %q", req.Message, resp.Message)
	}
}
