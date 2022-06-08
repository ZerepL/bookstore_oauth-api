package access_token

import (
	"strings"

	internalErrors "github.com/ZerepL/bookstore_oauth-api/src/utils/errors"
)

type Repository interface {
	GetById(string) (*AccessToken, *internalErrors.RestErr)
	Create(AccessToken) *internalErrors.RestErr
	UpdateExpirationTime(AccessToken) *internalErrors.RestErr
}

type Service interface {
	GetById(string) (*AccessToken, *internalErrors.RestErr)
	Create(AccessToken) *internalErrors.RestErr
	UpdateExpirationTime(AccessToken) *internalErrors.RestErr
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *internalErrors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, internalErrors.NewBadRequestError("invalid access token id")
	}

	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(at AccessToken) *internalErrors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.Create(at)
}

func (s *service) UpdateExpirationTime(at AccessToken) *internalErrors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(at)
}
