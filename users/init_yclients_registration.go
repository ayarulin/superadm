package users

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
)

type YclientsRegistrationTokenClaims struct {
	Name        string `json:"foo"`
	CompanyName string
	jwt.RegisteredClaims
}

type InitYclientsRegistration struct {
	ExtCompanyId string
	Name         string
	CompanyName  string
	Email        string
	PhoneNumber  string
	Signature    string
}

type RegistrationToken string

func(c InitYclientsRegistration) Handler(cmd *command) handlerFunc {
  return func (ctx context.Context) (any, error) {
    return "", nil
  }
}
