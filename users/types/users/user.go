package users

import (
	"strings"

	"github.com/DiptoChakrabarty/go-mvcs/users/utils/resterrors"
)

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Created   string `json:"created"`
}

func (usr User) Validate() *resterrors.RestErr {
	usr.Email = strings.TrimSpace(strings.ToLower(usr.Email))
	if usr.Email == "" {
		return resterrors.BadRequestError("invalid email address")
	}
	return nil
}
