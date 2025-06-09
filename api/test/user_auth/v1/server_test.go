package user_auth_test

import (
	"context"
	"testing"
	"time"

	userAuthPb "github.com/KonnorFrik/ChatServer/pkg/user_auth/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	// raddr = "localhost:9999"
	raddr = "0.0.0.0:9999"
)

var (
	userAuthClient userAuthPb.UserServiceClient
	baseCtx        = context.Background()
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

func TestCRUD(t *testing.T) {
	var (
		id       int64
		name     string = "user"
		email    string = "user@email.domail"
		password string = "password"
		role            = userAuthPb.Role_ROLE_USER
	)

	cases := []struct {
		name string
		f    func(*testing.T)
	}{

		{
			name: "Create user",
			f: func(t *testing.T) {
				resp, err := userAuthClient.Create(baseCtx, &userAuthPb.CreateUserRequest{
					Name:            name,
					Email:           email,
					Password:        password,
					PasswordConfirm: password,
					Role:            role,
				})
				if err != nil {
					t.Errorf("Got = %q\n", err)
				}
				id = resp.GetId()
			},
		},

		{
			name: "Get user",
			f: func(t *testing.T) {
				resp, err := userAuthClient.Get(baseCtx, &userAuthPb.GetUserRequest{
					Id: id,
				})
				if err != nil {
					t.Errorf("Got = %q\n", err)
				}
				if resp.GetId() != id {
					t.Errorf("Got = <%T, %[1]v>, Want = <%T, %[2]v\n", resp.GetId(), id)
				}
				if resp.GetName() != name {
					t.Errorf("Got = <%T, %[1]v>, Want = <%T, %[2]v\n", resp.GetName(), name)
				}
				if resp.GetEmail() != email {
					t.Errorf("Got = <%T, %[1]v>, Want = <%T, %[2]v\n", resp.GetEmail(), email)
				}
				if resp.GetRole() != role {
					t.Errorf("Got = <%T, %[1]v>, Want = <%T, %[2]v\n", resp.GetRole(), role)
				}
				var emptyTime time.Time
				if resp.GetCreatedAt().AsTime() == emptyTime {
					t.Errorf("Got = <%T, %[1]v>\n", resp.GetCreatedAt().AsTime(), role)
				}
				if resp.GetUpdatedAt().AsTime() == emptyTime {
					t.Errorf("Got = <%T, %[1]v>\n", resp.GetCreatedAt().AsTime(), role)
				}
			},
		},

        {
            name: "Update user name",
            f: func(t *testing.T) {
                var newName = "new name"
                _, err := userAuthClient.Update(baseCtx, &userAuthPb.UpdateUserRequest{
                    Id: id,
                    Name: &newName,
                })
                if err != nil {
                    t.Errorf("Got = %q\n", err)
                }
                resp, err := userAuthClient.Get(baseCtx, &userAuthPb.GetUserRequest{
                    Id: id,
                })
                if err != nil {
                    t.Errorf("Got = %q\n", err)
                }
                if resp.GetName() != newName {
					t.Errorf("Got = <%T, %[1]v>\n", resp.GetName(), newName)
                }
            },
        },

        {
            name: "Update user email",
            f: func(t *testing.T) {
                var newEmail = "newEmail@email.domain"
                _, err := userAuthClient.Update(baseCtx, &userAuthPb.UpdateUserRequest{
                    Id: id,
                    Email: &newEmail,
                })
                if err != nil {
                    t.Errorf("Got = %q\n", err)
                }
                resp, err := userAuthClient.Get(baseCtx, &userAuthPb.GetUserRequest{
                    Id: id,
                })
                if err != nil {
                    t.Errorf("Got = %q\n", err)
                }
                if resp.GetEmail() != newEmail {
					t.Errorf("Got = <%T, %[1]v>\n", resp.GetEmail(), newEmail)
                }
            },
        },

        {
            name: "Update user name and email",
            f: func(t *testing.T) {
                var newName = "ABOBA"
                var newEmail = "ABOBA@email.domain"
                _, err := userAuthClient.Update(baseCtx, &userAuthPb.UpdateUserRequest{
                    Id: id,
                    Name: &newName,
                    Email: &newEmail,
                })
                if err != nil {
                    t.Errorf("Got = %q\n", err)
                }
                resp, err := userAuthClient.Get(baseCtx, &userAuthPb.GetUserRequest{
                    Id: id,
                })
                if err != nil {
                    t.Errorf("Got = %q\n", err)
                }
                if resp.GetName() != newName {
					t.Errorf("Got = <%T, %[1]v>\n", resp.GetName(), newName)
                }
                if resp.GetEmail() != newEmail {
					t.Errorf("Got = <%T, %[1]v>\n", resp.GetEmail(), newEmail)
                }
            },
        },

        {
            name: "Delete user",
            f: func(t *testing.T) {
                _, err := userAuthClient.Delete(baseCtx,  &userAuthPb.DeleteUserRequest{
                    Id: id,
                })
                if err != nil {
                    t.Errorf("Got = %q\n", err)
                }
            },
        },
	}

	for _, tt := range cases{
        t.Run(tt.name, tt.f)
	}
}

// func TestCreate(t *testing.T) {
// 	cases := []struct {
// 		name      string
// 		reqData   *userAuthPb.CreateUserRequest
// 		wantResp  *userAuthPb.CreateUserResponse
// 		wantError error
// 	}{
// 		{
// 			"Simple",
// 			&userAuthPb.CreateUserRequest{
// 				Name:            "user",
// 				Email:           "user@email.domail",
// 				Password:        "password",
// 				PasswordConfirm: "password",
// 				Role:            userAuthPb.Role_ROLE_USER,
// 			},
// 			&userAuthPb.CreateUserResponse{},
// 			nil,
// 		},
// 	}
//
// 	for _, tt := range cases {
// 		t.Run(tt.name, func(t *testing.T) {
// 			resp, err := userAuthClient.Create(baseCtx, tt.reqData)
//
// 			if err != tt.wantError {
// 				t.Fatalf("Got = %q, Want = %q\n", err, tt.wantError)
// 			}
//
// 			_ = resp.Id
// 		})
// 	}
// }
