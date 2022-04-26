package models

import (
	"strings"

	"github.com/DiptoChakrabarty/go-mvcs/users/utils/resterrors"
)

type User struct {
	ID        uint64 `gorm:"primaryKey;auto_increment" json:"id"`
	FirstName string `json:"first" binding:"required" gorm:"type:varchar(10)"`
	LastName  string `json:"last" binding:"required" gorm:"type:varchar(10)"`
	Email     string `json:"email" binding:"required" gorm:"type:varchar(10)"`
	Created   string `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (usr User) Validate() *resterrors.RestErr {
	usr.Email = strings.TrimSpace(strings.ToLower(usr.Email))
	if usr.Email == "" {
		return resterrors.BadRequestError("invalid email address")
	}
	return nil
}
