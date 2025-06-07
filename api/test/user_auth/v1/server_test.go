package user_auth_test

import (
	"context"
	"testing"

	userAuthPb "github.com/KonnorFrik/ChatServer/pkg/user_auth/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
    raddr = "localhost:9999"
)

var (
    userAuthClient userAuthPb.UserServiceClient
    baseCtx = context.Background()
)

func init() {
    opts := []grpc.DialOption{
        grpc.WithTransportCredentials(insecure.NewCredentials()),
    }

    conn, err := grpc.NewClient(raddr, opts...)

    if err != nil {
        panic(err)
    }

    userAuthClient = userAuthPb.NewUserServiceClient(conn)
}

func TestCreate(t *testing.T) {
    cases := []struct{
        name string
        reqData *userAuthPb.CreateUserRequest
        wantResp *userAuthPb.CreateUserResponse
        wantError error

    }{
        {
            "Simple",
            &userAuthPb.CreateUserRequest{
                Name: "user",
                Email: "user@email.domail",
                Password: "password",
                PasswordConfirm: "password",
                Role: userAuthPb.Role_ROLE_USER,
            },
            &userAuthPb.CreateUserResponse{},
            nil,
        },
    }

    for _, tt := range cases {
        t.Run(tt.name, func(t *testing.T) {
            resp, err := userAuthClient.Create(baseCtx, tt.reqData)

            if err != tt.wantError {
                t.Fatalf("Got = %q, Want = %q\n", err, tt.wantError)
            }

            _ = resp.Id
        })
    }
}
