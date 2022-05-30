package database

import (
	"github.com/DiptoChakrabarty/go-mvcs/logger"
	"github.com/DiptoChakrabarty/go-mvcs/oauth/access_token"
	"github.com/DiptoChakrabarty/go-mvcs/oauth/client/cassandra"
	"github.com/DiptoChakrabarty/go-mvcs/users/utils/resterrors"
)

type DBRepository interface {
	GetById(string) (*access_token.AccessToken, *resterrors.RestErr)
}

type dbrepository struct{}

func NewDBRepository() DBRepository {
	return dbrepository{}
}

func (db dbrepository) GetById(string) (*access_token.AccessToken, *resterrors.RestErr) {
	session, err := cassandra.GetDBSession()
	if err != nil {
		logger.Error("Error while creating DB Session", err)
	}
	defer session.Close()
	return nil, nil
}
