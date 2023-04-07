package models

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/anyingiit/GoReactResourceManagement/structs"
	"github.com/anyingiit/GoReactResourceManagement/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username           string
	Password           string
	Salt               string
	MustChangePassword bool //要求登录时必须更改密码
	Name               string
	Age                int
	RoleID             uint
	Role               Role `gorm:"foreignKey:RoleID"`
}

var (
	ErrUserMustChangePassword = fmt.Errorf("must change password")
)

// new user
func NewUser(username, password, name string, age int, role Role) (*User, error) {
	user := &User{
		Name: name,
		Age:  age,
		Role: role,
	}
	// the sale will be generated in SetPassword function
	if err := user.SetPassword(password); err != nil {
		return nil, err
	}
	if err := user.SetUsername(username); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) SetPassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return err
	}
	hash := sha256.Sum256([]byte(password + hex.EncodeToString(salt)))
	u.Password = hex.EncodeToString(hash[:])
	u.Salt = hex.EncodeToString(salt)
	return nil
}

func (u *User) VaildPassword(password string) bool {
	hash := sha256.Sum256([]byte(password + u.Salt))
	return u.Password == hex.EncodeToString(hash[:])
}

// check user is must change password
func (u *User) CheckMustChangePassword() bool {
	return u.MustChangePassword
}

// check user is valid
func (u *User) CheckValid() error {
	if u.CheckMustChangePassword() {
		return ErrUserMustChangePassword
	}
	return nil
}

func (u *User) SetUsername(username string) error {
	if len(username) < 4 {
		return errors.New("username must be at least 4 characters")
	}
	u.Username = username
	return nil
}

// set must change password
func (u *User) SetMustChangePassword(mustChangePassword bool) {
	u.MustChangePassword = mustChangePassword
}

// generate token
func (u *User) GenerateToken() (string, error) {
	//TODO 应当读取config
	return utils.GenerateToken(structs.TokenData{
		UserId:  u.ID,
		TokenId: uuid.NewString(), // TODO: 未来可以考虑记录下来，以便在用户注销时，可以将其从数据库中删除
	}, "main service", 24*7*time.Hour, "anyingiit")
}
