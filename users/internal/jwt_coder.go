package internal

import (
	"time"

	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwt"
	"superadmin.ru/users/internal/core"
)

type JwtCoder struct {
	secret []byte
}

func NewJwtCoder(secret string) *JwtCoder {
	return &JwtCoder{secret: []byte(secret)}
}

func (j *JwtCoder) EncodeYclientsRegistration(reg *core.YclientsRegistration) (string, error) {
	token, err := jwt.NewBuilder().
		Subject(reg.ExtCompanyId).
		Claim("name", reg.Name).
		Claim("company_name", reg.CompanyName).
		Claim("email", reg.Email).
		Claim("phone_number", reg.PhoneNumber).
		Expiration(time.Now().Add(time.Minute * 15)).
		Build()

	if err != nil {
		return "", err
	}

	signedToken, err := jwt.Sign(token, jwt.WithKey(jwa.HS256(), j.secret))

	if err != nil {
		return "", err
	}

	return string(signedToken), nil
}

func (j *JwtCoder) DecodeYclientsRegistration(signedString string) (*core.YclientsRegistration, error) {
  var (
    name string
    company_id string
    company_name string
    email string
    phone_number string
  )

	token, err := jwt.ParseString(signedString, jwt.WithKey(jwa.HS256(), j.secret))

	if err != nil {
		return nil, err
	}

  token.Get("subject", &company_id)
  token.Get("name", &name)
  token.Get("company_name", &company_name)
  token.Get("email", &email)
  token.Get("phone_number", &phone_number)

	return core.NewYclientsRegistration(
		company_id,
    name,
    company_name,
    email,
    phone_number,
	)
}
