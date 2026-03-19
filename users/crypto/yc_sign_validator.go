package crypto

import  "superadmin.ru/pkg/crypto"

type YcSignValidator struct {
  userToken string
}

func NewYcSignValidator(userToken string) *YcSignValidator{
  return &YcSignValidator{
    userToken: userToken,
  }
}

func(v *YcSignValidator) ValidateSign(data, sign string) bool { 
  return crypto.ValidateSign(data, sign, v.userToken)
}

