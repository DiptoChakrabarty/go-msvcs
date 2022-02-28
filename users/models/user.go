package models

import (
	"strings"

	"github.com/DiptoChakrabarty/go-mvcs/users/utils/resterrors"
)

type User struct {
	Id        int64  `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string `json:"first_name" binding:"required" gorm:"type:varchar(10)"`
	LastName  string `json:"last_name" binding:"required" gorm:"type:varchar(10)"`
	Email     string `json:"email" binding:"required" gorm:"type:varchar(10)"`
	Created   string `json:"created" gorm:"default:CURRENT_TIMESTAMP"`
}

func (usr User) Validate() *resterrors.RestErr {
	usr.Email = strings.TrimSpace(strings.ToLower(usr.Email))
	if usr.Email == "" {
		return resterrors.BadRequestError("invalid email address")
	}
	return nil
}
