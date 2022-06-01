package access_token_service

import (
	"strings"

	"github.com/DiptoChakrabarty/go-mvcs/oauth/access_token"
	"github.com/DiptoChakrabarty/go-mvcs/resterrors"
)

type Repository interface {
	GetById(string) (*access_token.AccessToken, resterrors.RestErr)
}

type AccessTokenService interface {
	GetById(string) (*access_token.AccessToken, resterrors.RestErr)
}

type service struct {
	repository Repository
}

func NewAccessTokenService(repo Repository) AccessTokenService {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(accessTokenID string) (*access_token.AccessToken, resterrors.RestErr) {
	accessTokenID = strings.TrimSpace(accessTokenID)
	if len(accessTokenID) == 0 {
		return nil, resterrors.BadRequestError("Invalid AccessToken given")
	}
	accessToken, err := s.repository.GetById(accessTokenID)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
