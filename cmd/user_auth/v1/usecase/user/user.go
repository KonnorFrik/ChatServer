/*
Flow:

	gRPC Request -> User -> DB user -> User -> gRPC Response
*/
package user

import (
	"github.com/KonnorFrik/ChatServer/pkg/sql/models"
	userAuthPb "github.com/KonnorFrik/ChatServer/pkg/user_auth/v1"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
)

/*
User - data transfer object used for convert data from/to request/DataBase_model through this struct
From/To methods implements fluid interface
From/To methods is too dumb and any validation or transformation must be before or after call them
*/
type User struct {
    ID int64
    Name string
    Email string
    Password string
    Role string
}

// FromGrpcRequest - Just copy info from 'req' into 'u'.
func (u *User) FromGrpcRequest(req *userAuthPb.CreateUserRequest) *User {
    u.Name = req.GetName()
    u.Email = req.GetEmail()
    u.Password = req.GetPassword()
    u.Role = userAuthPb.Role_name[int32(req.GetRole())]
    return u
}

// ToGrpcResponse - Just copy info from 'u' into 'resp'.
func (u *User) ToGrpcCreateResponse(resp *userAuthPb.CreateUserResponse) *User {
    resp.Id = u.ID
    return u
}

// FromDbModel - Just copy info from 'model' into 'u'.
func (u *User) FromDbModel(model *models.User) *User {
    u.ID = model.ID
    u.Name = model.Name
    u.Email = model.Email
    u.Password = model.Password
    return u
}

// FromDbModel - Just copy info from 'u' into 'model'.
func (u *User) ToDbCreateParams(model *models.CreateUserParams) *User {
    model.Name = u.Name
    model.Email = u.Email
    model.Password = u.Password
    model.Role = pgtype.Int4{Int32: userAuthPb.Role_value[u.Role], Valid: true}
    return u
}

// IsValid - check all fields of 'u' and report is they valid.
func (u *User) IsValid() bool {
    // validation stage 1: check is data exist
    if u.Role == "" || u.Password == "" || u.Name == "" || u.Email == "" {
        return false
    }

    return true
}

// HashPassword - hash the password inside 'u' and overwrite it with hashed.
func (u *User) HashPassword() error {
    passwordCrypted, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

    if err != nil {
        return err
    }
    
    u.Password = string(passwordCrypted)
    return nil
}

// ComparePassword - compare plain password inside 'u' with given hashed 'hashed' password.
func (u *User) ComparePassword(hashed string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(u.Password))
}
