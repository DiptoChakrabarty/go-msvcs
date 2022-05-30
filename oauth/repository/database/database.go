package database

import (
	"github.com/DiptoChakrabarty/go-mvcs/logger"
	"github.com/DiptoChakrabarty/go-mvcs/oauth/access_token"
	"github.com/DiptoChakrabarty/go-mvcs/oauth/client/cassandra"
	"github.com/DiptoChakrabarty/go-mvcs/users/utils/resterrors"
)

const (
	queryGetAccessToken = "SELECT access_token, user_id, client_id, expires FROM access_token WHERE access_token=?"
)

type DBRepository interface {
	GetById(string) (*access_token.AccessToken, *resterrors.RestErr)
}

type dbrepository struct{}

func NewDBRepository() DBRepository {
	return dbrepository{}
}

func (db dbrepository) GetById(id string) (*access_token.AccessToken, *resterrors.RestErr) {
	var result access_token.AccessToken
	session, err := cassandra.GetDBSession()
	if err != nil {
		logger.Error("Error while creating DB Session", err)
		return nil, resterrors.BadRequestError(err.Error())
	}
	defer session.Close()

	if err := session.Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires,
	); err != nil {
		logger.Error("Unable to retreive ID", err)
		return nil, resterrors.BadRequestError(err.Error())
	}
	return &result, nil
}
