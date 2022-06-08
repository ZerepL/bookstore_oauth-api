// @title           Bookstore OAuth API
// @version         1.0
// @description     API to work as security layer, generating and validating tokens.
// @termsOfService  http://swagger.io/terms/

// @host      localhost:8080
// @BasePath  /oauth

// @securityDefinitions.basic  BasicAuth
package main

import (
	"github.com/ZerepL/bookstore_oauth-api/src/app"
)

func main() {
	app.StartApplication()
}
