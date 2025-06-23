package integration

import (
	"context"
	"net"

	"github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/internal/handler"
	"github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/internal/server"
	v1 "github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/pkg/proto/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

// dialer spins up an in-memory gRPC server (Echo, Ping, Health, Reflection)
// and returns a client connection plus a cleanup functor.
func dialer() (*grpc.ClientConn, func(), error) {
	lis := bufconn.Listen(bufSize)

	grpcServer := grpc.NewServer()

	// Register business services
	v1.RegisterEchoServiceServer(grpcServer, handler.NewEchoServiceServer())
	v1.RegisterPingServiceServer(grpcServer, handler.NewPingServiceServer())

	// Register health + reflection
	grpc_health_v1.RegisterHealthServer(grpcServer, server.NewHealthServer())
	reflection.Register(grpcServer)

	// Start serving
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			panic("bufconn Serve error: " + err.Error())
		}
	}()

	// Build client conn
	ctx := context.Background()
	conn, err := grpc.DialContext(
		ctx,
		"bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) { return lis.Dial() }),
		grpc.WithInsecure(),
	)
	if err != nil {
		lis.Close()
		grpcServer.Stop()
		return nil, nil, err
	}

	cleanup := func() {
		conn.Close()
		grpcServer.GracefulStop()
		lis.Close()
	}

	return conn, cleanup, nil
}
