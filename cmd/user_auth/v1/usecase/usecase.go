/*
Business logic of user_auth/v1 service
Work flow:

	convert request into user
	do job with user and db
*/
package usecase

import (
	"context"

	"github.com/KonnorFrik/ChatServer/cmd/user_auth/v1/usecase/db"
	"github.com/KonnorFrik/ChatServer/cmd/user_auth/v1/usecase/user"

	userAuthPb "github.com/KonnorFrik/ChatServer/pkg/user_auth/v1"
)

func Create(ctx context.Context, req *userAuthPb.CreateUserRequest) (*user.User, error) {
    var u = new(user.User)
    err := u.FromGrpcRequest(req)

    if err != nil {
        return nil, err
    }

    db.DB().Queries.CreateUser()

    return u, nil
}

func Get(ctx context.Context, req *userAuthPb.GetUserRequest) (*user.User, error) {

    return nil, nil
}

func Update(ctx context.Context, req *userAuthPb.UpdateUserRequest) (*user.User, error) {
    
    return nil, nil
}

func Delete(ctx context.Context, req *userAuthPb.DeleteUserRequest) error {
    return nil
}

