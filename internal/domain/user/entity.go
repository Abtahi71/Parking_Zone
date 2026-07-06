package user

import (
	"gotickets/internal/domain/user/types"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Role     types.Role
}

func (u *User) hashPassword(password string)error{
	hash,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err!=nil{
		return err
	}
	u.Password = string(hash)
	return nil
}

func (u *User) checkPassword(password string)error{
	return bcrypt.CompareHashAndPassword(([]byte(u.Password)),[]byte(password))
}