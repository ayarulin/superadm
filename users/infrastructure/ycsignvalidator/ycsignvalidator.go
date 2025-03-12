package ycsignvalidator

import  "superadmin.ru/pkg/crypto"

type validator struct {
  userToken string
}

func New(userToken string) *validator{
  return &validator{
    userToken: userToken,
  }
}

func(v *validator) ValidateSign(data, sign string) bool { 
  return crypto.ValidateSign(data, sign, v.userToken)
}

