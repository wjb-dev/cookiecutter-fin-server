package handler_test

import (
	"context"
	"github.com/google/go-cmp/cmp/cmpopts"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/internal/handler"
	v1 "github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/pkg/proto/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestPing(t *testing.T) {
	srv := handler.NewPingServiceServer()

	want := &v1.PingResponse{
		Service: "{{ cookiecutter.project_slug }}",
		Status:  "SERVING",
		Version: "0.1.0",
	}

	resp, err := srv.Ping(context.Background(), &emptypb.Empty{})
	if err != nil {
		t.Fatalf("Ping() returned error: %v", err)
	}

	if diff := cmp.Diff(want, resp, cmpopts.IgnoreUnexported(v1.PingResponse{})); diff != "" {
		t.Errorf("Ping() response mismatch (-want +got):\n%s", diff)
	}
}
