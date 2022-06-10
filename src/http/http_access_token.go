package http

import (
	"github.com/ZerepL/bookstore_oauth-api/src/services/access_token_service"
	internalErrors "github.com/ZerepL/bookstore_utils/internal_errors"
	"net/http"

	atDomain "github.com/ZerepL/bookstore_oauth-api/src/domain/access_token"
	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
}

type accessTokenHandler struct {
	service access_token_service.Service
}

func NewAccessTokenHandler(service access_token_service.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

// ShowAccount godoc
// @Summary      Get token
// @Description  get token info
// @Tags         oauth
// @Produce      json
// @Param        access_token_id   path      string  true  "token name"
// @Success      200  {object}  access_token.AccessToken
// @Failure      400  {object}  internalErrors.RestErr
// @Failure      404  {object}  internalErrors.RestErr
// @Failure      500  {object}  internalErrors.RestErr
// @Router       /oauth/{access_token_id} [get]
func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessToken, err := handler.service.GetById(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

// ShowAccount godoc
// @Summary      Create Token
// @Description  create a token user
// @Tags         oauth
// @Produce      json
// @Success      200  {object}  access_token.AccessToken
// @Failure      400  {object}  internalErrors.RestErr
// @Failure      500  {object}  internalErrors.RestErr
// @Router       /oauth [post]
func (handler *accessTokenHandler) Create(c *gin.Context) {
	var request atDomain.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := internalErrors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	accessToken, err := handler.service.Create(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}
