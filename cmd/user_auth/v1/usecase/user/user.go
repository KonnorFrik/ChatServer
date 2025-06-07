/*
Flow:
    gRPC Request -> User -> DB user -> User -> gRPC Response
*/
package user

import (
    "golang.org/x/crypto/bcrypt"
	userAuthPb "github.com/KonnorFrik/ChatServer/pkg/user_auth/v1"
	"github.com/KonnorFrik/ChatServer/pkg/sql/models"
)

/*
User - data transfer object used for convert data from/to request/DataBase_model through this struct
From/To methods implements fluid interface
From/To methods is too dumb and any validation or transformation must be before or after call them
*/
type User struct {
    ID int64
    Name string
    Password string
    Role userAuthPb.Role
}

// FromGrpcRequest - Just copy info from 'req' into 'u'
func (u *User) FromGrpcRequest(req *userAuthPb.CreateUserRequest) *User {

    return u
}

// ToGrpcResponse - Just copy info from 'u' into 'resp'
func (u *User) ToGrpcResponse(resp *userAuthPb.CreateUserResponse) *User {

    return u
}

// FromDbModel - Just copy info from 'model' into 'u'
func (u *User) FromDbModel(model *models.User) *User {

    return u
}

// FromDbModel - Just copy info from 'u' into 'model'
func (u *User) ToDbCreateParams(model *models.CreateUserParams) *User {
    
    return u
}

func (u *User) HashPassword() error {
    var passwordCrypted []byte
    var err error
    passwordCrypted, err = bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

    if err != nil {
        return err
    }
    
    u.Password = string(passwordCrypted)
    return nil
}

func (u *User) ComparePassword(hashed string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(u.Password))
}
