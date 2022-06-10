package database

import (
	"github.com/DiptoChakrabarty/go-mvcs/logger"
	"github.com/DiptoChakrabarty/go-mvcs/oauth/access_token"
	"github.com/DiptoChakrabarty/go-mvcs/oauth/client/cassandra"
	"github.com/DiptoChakrabarty/go-mvcs/resterrors"
)

const (
	queryGetAccessToken   = "SELECT access_token, user_id, client_id, expires FROM access_token WHERE access_token=?"
	queryCreateToken      = "INSERT INTO access_token(access_token, user_id, client_id expires) VALUES (?,?,?,?);"
	queryUpdateExpiryTime = "UPDATE access_token SET expires=? WHERE access_token=?;"
)

type DBRepository interface {
	GetById(string) (*access_token.AccessToken, resterrors.RestErr)
	Create(access_token.AccessToken) resterrors.RestErr
	UpdateExpiryTime(access_token.AccessToken) resterrors.RestErr
}

type dbrepository struct{}

func NewDBRepository() DBRepository {
	return dbrepository{}
}

func (db dbrepository) GetById(id string) (*access_token.AccessToken, resterrors.RestErr) {
	var result access_token.AccessToken
	if err := cassandra.GetDBSession().Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires,
	); err != nil {
		logger.Error("Unable to retreive ID", err)
		return nil, resterrors.NotFound("Id is not found")
	}
	return &result, nil
}

func (db dbrepository) Create(at access_token.AccessToken) resterrors.RestErr {
	if err := cassandra.GetDBSession().Query(queryCreateToken,
		at.AccessToken,
		at.UserId,
		at.ClientId,
		at.Expires).Exec(); err != nil {
		logger.Error("Unable to create the access token", err)
		return resterrors.InternalServerError("Unable to create access token", err)
	}
	return nil
}

func (db dbrepository) UpdateExpiryTime(at access_token.AccessToken) resterrors.RestErr {
	if err := cassandra.GetDBSession().Query(queryUpdateExpiryTime,
		at.Expires,
		at.AccessToken).Exec(); err != nil {
		logger.Error("Unable to update access token expiry time", err)
		return resterrors.InternalServerError("Unable to update access token expiry time", err)
	}
	return nil
}
