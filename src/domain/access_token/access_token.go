package access_token

import (
	internalErrors "github.com/ZerepL/bookstore_oauth-api/src/utils/errors"
	"strings"
	"time"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken() *AccessToken {
	return &AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Validate() *internalErrors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return internalErrors.NewBadRequestError("invalid access token id")
	}

	if at.UserId <= 0 {
		return internalErrors.NewBadRequestError("invalid user id")
	}

	if at.ClientId <= 0 {
		return internalErrors.NewBadRequestError("invalid client id")
	}

	if at.Expires <= 0 {
		return internalErrors.NewBadRequestError("invalid expiration time")
	}

	return nil
}
