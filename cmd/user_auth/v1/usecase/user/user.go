package user

import (
	userAuthPb "github.com/KonnorFrik/ChatServer/pkg/user_auth/v1"
)

type User struct {
    ID int64
    Name string
    Password string
    Role userAuthPb.Role
}

// FromRequest - copy info from req into u
func (u *User) FromRequest(req *userAuthPb.CreateUserRequest) error {

    return nil
}

func (u *User) ToResponse() (*userAuthPb.CreateUserResponse, error) {

    return nil, nil
}
