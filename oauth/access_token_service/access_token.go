package access_token_service

import (
	"strings"

	"github.com/DiptoChakrabarty/go-mvcs/logger"
	"github.com/DiptoChakrabarty/go-mvcs/oauth/access_token"
	"github.com/DiptoChakrabarty/go-mvcs/resterrors"
)

type Repository interface {
	GetById(string) (*access_token.AccessToken, resterrors.RestErr)
	Create(access_token.AccessToken) resterrors.RestErr
	UpdateExpiryTime(access_token.AccessToken) resterrors.RestErr
}

type AccessTokenService interface {
	GetById(string) (*access_token.AccessToken, resterrors.RestErr)
	Create(access_token.AccessToken) resterrors.RestErr
	UpdateExpiryTime(access_token.AccessToken) resterrors.RestErr
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
		logger.Error("unable to validate token while obtaining token", err)
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(at access_token.AccessToken) resterrors.RestErr {
	if err := at.ValidateToken(); err != nil {
		logger.Error("unable to validate token while creating", err)
		return err
	}
	return s.repository.Create(at)
}

func (s *service) UpdateExpiryTime(at access_token.AccessToken) resterrors.RestErr {
	if err := at.ValidateToken(); err != nil {
		logger.Error("unable to validate token while updating", err)
		return err
	}
	return s.repository.UpdateExpiryTime(at)
}
