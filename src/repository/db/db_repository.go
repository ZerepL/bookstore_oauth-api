package db

import (
	"github.com/ZerepL/bookstore_oauth-api/src/domain/access_token"
	internalErrors "github.com/ZerepL/bookstore_oauth-api/src/utils/errors"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *internalErrors.RestErr)
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

type dbRepository struct {
}

func (r *dbRepository) GetById(string) (*access_token.AccessToken, *internalErrors.RestErr) {
	return nil, internalErrors.NewInternalServerError("Not implemented")
}
