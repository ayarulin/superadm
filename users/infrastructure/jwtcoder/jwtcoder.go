package jwtcoder

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	core "superadmin.ru/users/internal"
)

type jwtcoder struct {
	secret []byte
}

func New(secret string) *jwtcoder {
	return &jwtcoder{secret: []byte(secret)}
}

func (j *jwtcoder) EncodeYclientsRegistration(reg *core.YclientsRegistration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	//
	// // Sign and get the complete encoded token as a string using the secret

	return token.SignedString(j.secret)
	// // generate confirmation code
	// // generate jwt with all reg data + code

}

func (j *jwtcoder) DecodeYclientsRegistration(token string) (*core.YclientsRegistration, error) {
	return nil, nil
}
