package services

import (
	"github.com/DiptoChakrabarty/go-mvcs/users/models"
	"github.com/DiptoChakrabarty/go-mvcs/users/utils/resterrors"
)

type UserService struct {
	model models.Model
}

func (svc *UserService) AddUser(u models.User) (*models.User, *resterrors.RestErr) {
	if err := u.Validate(); err != nil {
		return nil, err
	}
	err := svc.model.Save(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (svc *UserService) GetUser(id uint64) (*models.User, *resterrors.RestErr) {
	u, err := svc.model.Find(id)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (svc *UserService) UpdateUser(u models.User) (*models.User, *resterrors.RestErr) {
	if err := u.Validate(); err != nil {
		return nil, err
	}
	err := svc.model.Update(u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (svc *UserService) DeleteUser(id uint64) *resterrors.RestErr {
	err := svc.model.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
