package rest

import (
	"encoding/json"
	"errors"
	"github.com/ZerepL/bookstore_oauth-api/src/domain/users"
	internalErrors "github.com/ZerepL/bookstore_utils/internal_errors"
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

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, internalErrors.RestErr)
}

type usersRepository struct{}

func NewRestUsersRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, internalErrors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	response := userRestClient.Post(userLogin, request)

	if response == nil || response.Response == nil {
		return nil, internalErrors.NewInternalServerError("invalid restclient response when trying to login user", errors.New("restclient error"))
	}

	if response.StatusCode > 299 {
		apiErr, err := internalErrors.NewRestErrorFromBytes(response.Bytes())
		if err != nil {
			return nil, internalErrors.NewInternalServerError("invalid error interface when trying to login user", err)
		}
		return nil, apiErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, internalErrors.NewInternalServerError("error when trying to unmarshal users login response", errors.New("json parsing error"))
	}
	return &user, nil
}
