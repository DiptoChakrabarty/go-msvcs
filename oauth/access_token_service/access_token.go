package service

import (
	"github.com/DiptoChakrabarty/go-mvcs/oauth/access_token"
	"github.com/DiptoChakrabarty/go-mvcs/users/utils/resterrors"
)

type Repository interface {
	GetById(string) (*access_token.AccessToken, *resterrors.RestErr)
}

type AccessTokenService interface {
	GetById(string) (*access_token.AccessToken, *resterrors.RestErr)
}

type service struct {
	repository Repository
}

func NewAccessTokenService(repo Repository) AccessTokenService {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(string) (*access_token.AccessToken, *resterrors.RestErr) {
	return nil, nil
}
