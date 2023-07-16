package controller_auth

import (
	"errors"
	server_config "github.com/ESPOIR-DITE/tim-api/config/server.config"
)

func IsAllowed(app *server_config.Env, token string) (*string, error) {
	if token == "" {
		return nil, errors.New("You are not allowed to access this resource")
	}
	jwtClaim, err := app.SecurityService.ValidateToken(token)
	if err != nil {
		return nil, errors.New("Please login again. You have an invalid token")
	}
	return &jwtClaim.Id, nil
}
