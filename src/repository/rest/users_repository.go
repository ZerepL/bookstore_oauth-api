package rest

import (
	"encoding/json"
	"github.com/ZerepL/bookstore_oauth-api/src/domain/users"
	internalErrors "github.com/ZerepL/bookstore_oauth-api/src/utils/errors"
	"github.com/federicoleon/golang-restclient/rest"
	"os"
	"time"
)

const (
	bookstoreUsersEndpoint = "bookstore_users_endpoint"
	userLogin              = "/users/login"
)

var (
	bookstoreUsersEndPoint = os.Getenv(bookstoreUsersEndpoint)
	userRestClient         = rest.RequestBuilder{
		BaseURL: bookstoreUsersEndPoint,
		Timeout: 100 * time.Millisecond,
	}
)

type RestUserRepository interface {
	LoginUser(string, string) (*users.User, *internalErrors.RestErr)
}

type usersRepository struct{}

func NewRepository() RestUserRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, *internalErrors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	response := userRestClient.Post(userLogin, request)
	if response == nil || response.Response == nil {
		return nil, internalErrors.NewInternalServerError("invalid restClient response when trying to login user")
	}
	if response.StatusCode > 299 {
		var restErr internalErrors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, internalErrors.NewInternalServerError("invalid error interface when trying to login user")
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, internalErrors.NewInternalServerError("error when trying to unmarshal users login response")
	}
	return &user, nil
}
