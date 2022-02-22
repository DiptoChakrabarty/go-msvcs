package services

import (
	"github.com/DiptoChakrabarty/go-mvcs/users/types/users"
	"github.com/DiptoChakrabarty/go-mvcs/users/utils/resterrors"
)

func AddUser(u users.User) (*users.User, *resterrors.RestErr) {
	if err := u.Validate(); err != nil {
		return nil, err
	}
	return nil, nil
}

func GetUser(id uint64) (*users.User, *resterrors.RestErr) {
	return nil, nil
}

func DeleteUser(id uint64) (*users.User, *resterrors.RestErr) {
	return nil, nil
}
