/*
Business logic of user_auth/v1 service
Work flow:
    convert request into user 
    do job with user and db
*/
package usecase

import (
    "context"
	userAuthPb "github.com/KonnorFrik/ChatServer/pkg/user_auth/v1"
	"github.com/KonnorFrik/ChatServer/cmd/user_auth/v1/usecase/user"
)

func Create(ctx context.Context, req *userAuthPb.CreateUserRequest) (*user.User, error) {

    return nil, nil
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

