/*
Simple gRPC user_auth server implemented user_auth/v1
! Used in-memory storage
*/
package main

import (
	"context"
	"net"

	recoverInterceptor "github.com/KonnorFrik/ChatServer/pkg/interceptor/recover"
	logging "github.com/KonnorFrik/ChatServer/pkg/logging"
	userAuthPb "github.com/KonnorFrik/ChatServer/pkg/user_auth/v1"
	"google.golang.org/grpc"
)

type server struct {
    userAuthPb.UnimplementedUserServiceServer
}

var (
    logger = logging.New()
)

const (
    laddr = ":9999"
)

func main() {
    listener, err := net.Listen("tcp", laddr)

    if err != nil {
        logger.Error("Listen", "error", err)
        return
    }

    userServer := &server{}
    grpcServer := grpc.NewServer(
        grpc.ChainUnaryInterceptor(
            logger.UnaryServerInterceptor,
            recoverInterceptor.UnaryServerRecoverInterceptor,
        ),
    )
    userAuthPb.RegisterUserServiceServer(grpcServer, userServer)
    logger.Info("Listen at", "local address", laddr)

    err = grpcServer.Serve(listener)

    if err != nil {
        logger.Error("Serve", "error", err)
        return
    }
}
func (s *server) Create(ctx context.Context, req *userAuthPb.CreateUserRequest) (*userAuthPb.CreateUserResponse, error) {
    panic("TEST PANIC")
}
