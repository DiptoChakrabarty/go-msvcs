package models

import (
	"strings"

	"github.com/DiptoChakrabarty/go-mvcs/resterrors"
)

type User struct {
	ID        uint64 `gorm:"primaryKey;auto_increment" json:"id"`
	FirstName string `json:"first" binding:"required" gorm:"type:varchar(20)"`
	LastName  string `json:"last" binding:"required" gorm:"type:varchar(20)"`
	Email     string `json:"email" binding:"required" gorm:"type:varchar(20)"`
	Created   string `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Password  string `json:"password" binding:"required" gorm:"type:varchar(20)"`
}

func (usr User) Validate() *resterrors.RestErr {
	usr.Email = strings.TrimSpace(strings.ToLower(usr.Email))
	if usr.Email == "" {
		return resterrors.BadRequestError("invalid email address")
	}

	usr.Password = strings.TrimSpace(usr.Password)
	if usr.Password == "" {
		return resterrors.BadRequestError("empty password provided")
	}
	return nil
}
