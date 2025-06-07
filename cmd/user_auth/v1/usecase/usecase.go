/*
Business logic of user_auth/v1 service
Work flow:

	convert request into user
	do job with user and db
*/
package usecase

import (
	"context"
	"log"

	"github.com/KonnorFrik/ChatServer/cmd/user_auth/v1/usecase/db"
	"github.com/KonnorFrik/ChatServer/cmd/user_auth/v1/usecase/user"

	"github.com/KonnorFrik/ChatServer/pkg/sql/models"
	userAuthPb "github.com/KonnorFrik/ChatServer/pkg/user_auth/v1"
)

func Create(ctx context.Context, req *userAuthPb.CreateUserRequest) (*user.User, error) {
    var u = new(user.User)
    var createParams models.CreateUserParams

    if !u.FromGrpcRequest(req).IsValid() {
        return nil, ErrInvalidData
    }

    if err := u.HashPassword(); err != nil {
        log.Println("[usecase/CreateUser]: Hash error:", err)
        return nil, ErrInvalidData
    }

    u.ToDbCreateParams(&createParams)
    userDB, err := db.DB().Queries.CreateUser(ctx, createParams)

    if err != nil {
        return nil, WrapError(db.DB().WrapError(err))
    }

    u.FromDbModel(userDB)
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

