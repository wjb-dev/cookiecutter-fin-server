package handler_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/internal/handler"
	v1 "github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/pkg/proto/v1"
)

func TestEcho(t *testing.T) {
	srv := handler.NewEchoServiceServer()

	testCases := []struct {
		name string
		msg  string
	}{
		{"simple", "hello"},
		{"empty", ""},
		{"unicode", "こんにちは"},
	}

	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			resp, err := srv.Echo(context.Background(), &v1.EchoRequest{Message: tc.msg})
			if err != nil {
				t.Fatalf("Echo() returned error: %v", err)
			}

			if diff := cmp.Diff(tc.msg, resp.GetMessage()); diff != "" {
				t.Errorf("Echo() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
