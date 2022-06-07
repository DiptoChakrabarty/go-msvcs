package access_token

import (
	"strings"
	"time"

	"github.com/DiptoChakrabarty/go-mvcs/resterrors"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int    `json:"user_id"`
	ClientId    int    `json:"user_id"`
	Expires     int64  `json:"expires"`
}

func GenerateAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) ValidateToken() resterrors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if len(at.AccessToken) == 0 {
		return resterrors.BadRequestError("invalid access token id")
	}
	if at.UserId <= 0 {
		return resterrors.BadRequestError("invalid user id")
	}
	if at.ClientId <= 0 {
		return resterrors.BadRequestError("invalid client id")
	}
	if at.Expires <= 0 {
		return resterrors.BadRequestError("invalid expiry time")
	}
	return nil
}

func (at AccessToken) CheckExpired() bool {
	return false
}
