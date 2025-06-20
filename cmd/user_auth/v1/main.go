/*
Simple gRPC user_auth server implemented user_auth/v1
*/
package main

import (
	"context"
	"errors"
	"net"

	"github.com/KonnorFrik/ChatServer/cmd/user_auth/v1/usecase"
	recoverInterceptor "github.com/KonnorFrik/ChatServer/pkg/interceptor/recover"
	logging "github.com/KonnorFrik/ChatServer/pkg/logging"
	userAuthPb "github.com/KonnorFrik/ChatServer/pkg/user_auth/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

type server struct {
	userAuthPb.UnimplementedUserServiceServer
}

var (
	logger = logging.New()
    tracer *tracesdk.TracerProvider
)

const (
	laddr = ":9999"
)

func main() {
    tracer, err := NewTracer("http://localhost:14268/api/traces", "server")

    if err != nil {
        logger.Error("[Server/NewTracer]", "error", err)
    }

    defer tracer.Shutdown(context.Background())
	listener, err := net.Listen("tcp", laddr)

	if err != nil {
		logger.Error("[Server/Listen]", "error", err)
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

// WrapError - wrap usecase error into gRPC error with codes
func WrapError(err error) error {
	var code = codes.Internal
	var msg string
	switch {
	case errors.Is(err, usecase.ErrDoesNotExist):
		code = codes.NotFound
	case errors.Is(err, usecase.ErrAlreadyExist):
		code = codes.AlreadyExists
	case errors.Is(err, usecase.ErrInvalidData):
		code = codes.InvalidArgument
	case errors.Is(err, usecase.ErrDbNoAccess):
		// default = Internal
	case errors.Is(err, usecase.ErrUnknown):
		// default = Internal
	}

	return status.Error(code, msg)
}

func (s *server) Create(
	ctx context.Context,
	req *userAuthPb.CreateUserRequest,
) (
	*userAuthPb.CreateUserResponse,
	error,
) {
    ctx, span := tracer.Tracer("Server").Start(ctx, "Create")
    defer span.End()

	var response userAuthPb.CreateUserResponse
	user, err := usecase.Create(ctx, req)

	if err != nil {
		return nil, WrapError(err)
	}

	user.ToGrpcCreateResponse(&response)
	return &response, status.Error(codes.OK, "")
}

func (s *server) Get(
    ctx context.Context,
    req *userAuthPb.GetUserRequest,
) (
    *userAuthPb.GetUserResponse,
    error,
) {
    ctx, span := tracer.Tracer("Server").Start(ctx, "Get")
    defer span.End()

    user, err := usecase.Get(ctx, req)

    if err != nil {
        return nil, WrapError(err)
    }

    var response userAuthPb.GetUserResponse
    user.ToGrpcGetResponse(&response)
    return &response, status.Error(codes.OK, "")
}

func (s *server) Update(
    ctx context.Context,
    req *userAuthPb.UpdateUserRequest,
) (
    *userAuthPb.UpdateUserResponse,
    error,
) {
    ctx, span := tracer.Tracer("Server").Start(ctx, "Update")
    defer span.End()

    err := usecase.Update(ctx, req)

    if err != nil {
        return nil, WrapError(err)
    }

    var response userAuthPb.UpdateUserResponse
    return &response, status.Error(codes.OK, "")
}

func (s *server) Delete(
    ctx context.Context,
    req *userAuthPb.DeleteUserRequest,
) (
    *userAuthPb.DeleteUserResponse,
    error,
) {
    ctx, span := tracer.Tracer("Server").Start(ctx, "Delete")
    defer span.End()

    err := usecase.Delete(ctx, req)

    if err != nil {
        return nil, WrapError(err)
    }

    var response userAuthPb.DeleteUserResponse
    return &response, status.Error(codes.OK, "")
}
