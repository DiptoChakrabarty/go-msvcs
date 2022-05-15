package service

import (
	"github.com/DiptoChakrabarty/go-mvcs/oauth/access_token"
	"github.com/DiptoChakrabarty/go-mvcs/users/utils/resterrors"
)

type Service interface {
	GetById(string) (*access_token.AccessToken, *resterrors.RestErr)
}

type service struct {
}

func NewAccessTokenService() Service {
	return &service{}
}

func (s *service) GetById(string) (*access_token.AccessToken, *resterrors.RestErr) {
	return nil, nil
}
