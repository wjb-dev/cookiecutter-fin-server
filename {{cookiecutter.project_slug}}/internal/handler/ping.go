package handler

import (
	"context"
	v1 "github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/pkg/proto/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PingServiceServer struct {
	v1.UnimplementedPingServiceServer
}

func NewPingServiceServer() *PingServiceServer {
	return &PingServiceServer{}
}

func (s *PingServiceServer) Ping(ctx context.Context, _ *emptypb.Empty) (*v1.PingResponse, error) {
	return &v1.PingResponse{
		Service: "{{ cookiecutter.project_slug }}",
		Status:  "SERVING",
		Version: "0.1.0",
	}, nil
}
