/*
Business logic of user_auth/v1 service
Work flow:
    gRPC request -> User -> DB params -> User -> gRPC response
*/
package usecase

import (
	"context"
	"log"
	"log/slog"

	"github.com/KonnorFrik/ChatServer/cmd/user_auth/v1/usecase/user"
	"github.com/KonnorFrik/ChatServer/pkg/db"

	"github.com/KonnorFrik/ChatServer/pkg/logging"
	"github.com/KonnorFrik/ChatServer/pkg/sql/models"
	userAuthPb "github.com/KonnorFrik/ChatServer/pkg/user_auth/v1"
)

var logger = logging.New()

func Create(ctx context.Context, req *userAuthPb.CreateUserRequest) (*user.User, error) {
    var u = new(user.User)
    var createParams models.CreateUserParams

    if !u.FromGrpcCreateRequest(req).IsValid() {
        return nil, ErrInvalidData
    }

    if err := u.HashPassword(); err != nil {
        log.Println("[usecase/CreateUser]: Hash error:", err)
        return nil, ErrInvalidData
    }

    u.ToDbCreateParams(&createParams)
    userDB, err := db.DB().CreateUser(ctx, createParams)

    if err != nil {
        er := WrapError(db.DB().WrapError(err))
        logger.LogAttrs(
            ctx,
            slog.LevelError,
            "[usecase/CreateUser]",
            slog.String("error", er.Error()),
        )
        return nil, er
    }

    return u.FromDbModel(userDB), nil
}

func Get(ctx context.Context, req *userAuthPb.GetUserRequest) (*user.User, error) {
    var u = new(user.User)
    u.FromGrpcGetRequest(req)
    userDB, err := db.DB().UserByID(ctx, u.ID)

    if err != nil {
        er := WrapError(db.DB().WrapError(err))
        logger.LogAttrs(
            ctx,
            slog.LevelError,
            "[usecase/GetUser]",
            slog.String("error", er.Error()),
        )
        return nil, er
    }

    return u.FromDbModel(userDB), nil
}

func Update(ctx context.Context, req *userAuthPb.UpdateUserRequest) error {
    var (
        u = new(user.User)
        err error
    )
    u.FromGrpcUpdateRequest(req)

    switch {
    case len(u.Name) > 0 && len(u.Email) > 0:
        var updParams models.UpdateUserNameEmailParams
        u.ToDbUpdateNameEmailParams(&updParams)
        err = db.DB().UpdateUserNameEmail(ctx, updParams)

    case len(u.Name) > 0:
        var updParams models.UpdateUserNameParams
        u.ToDbUpdateNameParams(&updParams)
        err = db.DB().UpdateUserName(ctx, updParams)

    case len(u.Email) > 0:
        var updParams models.UpdateUserEmailParams
        u.ToDbUpdateEmailParams(&updParams)
        err = db.DB().UpdateUserEmail(ctx, updParams)

    default:
        err = ErrDoesNotExist
    }

    if err != nil {
        er := WrapError(db.DB().WrapError(err))
        logger.LogAttrs(
            ctx,
            slog.LevelError,
            "[usecase/UpdateUser]",
            slog.String("error", er.Error()),
        )
        return er
    }

    return nil
}

func Delete(ctx context.Context, req *userAuthPb.DeleteUserRequest) error {
    err := db.DB().DeleteUser(ctx, req.Id)

    if err != nil {
        er := WrapError(db.DB().WrapError(err))
        logger.LogAttrs(
            ctx,
            slog.LevelError,
            "[usecase/DeleteUser]",
            slog.String("error", er.Error()),
        )
        return er
    }

    return nil
}

// TODO: delete may be idempotent - need catch error and if it like Already deleted - return nil
