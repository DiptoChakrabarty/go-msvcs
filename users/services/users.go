package services

import (
	"github.com/DiptoChakrabarty/go-mvcs/users/models"
	"github.com/DiptoChakrabarty/go-mvcs/users/types/users"
	"github.com/DiptoChakrabarty/go-mvcs/users/utils/resterrors"
)

type UserService struct {
	model models.Model
}

func (svc *UserService) AddUser(u users.User) (*users.User, *resterrors.RestErr) {
	if err := u.Validate(); err != nil {
		return nil, err
	}
	return nil, nil
}

func (svc *UserService) GetUser(id uint64) (*users.User, *resterrors.RestErr) {
	return nil, nil
}

func (svc *UserService) DeleteUser(id uint64) (*users.User, *resterrors.RestErr) {
	return nil, nil
}
