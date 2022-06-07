package db

import (
	"github.com/ZerepL/bookstore_oauth-api/src/clients/cassandra"
	"github.com/ZerepL/bookstore_oauth-api/src/domain/access_token"
	internalErrors "github.com/ZerepL/bookstore_oauth-api/src/utils/errors"
)

const (
	queryGetAccessToken = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *internalErrors.RestErr)
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *internalErrors.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		return nil, internalErrors.NewInternalServerError(err.Error())
	}
	defer session.Close()

	var result access_token.AccessToken
	if err = session.Query(queryGetAccessToken, id).Scan(&result.AccessToken, &result.UserId, &result.ClientId, &result.Expires); err != nil {
		return nil, internalErrors.NewInternalServerError(err.Error())
	}
	return &result, nil
}
