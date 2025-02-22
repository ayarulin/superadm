package core

import (
	"fmt"
	"time"
  "superadmin.ru/pkg/validators"
)

type ActiveIntegrationId string
type ConfirmationTime time.Time
type AccountId string
type ExtCompanyId string
type CompanyName string

func NewActiveIntegrationId(s string) (ActiveIntegrationId, error) {
  err := validators.String.IsNotBlank(s)
  err := validators.String.MaxLen(s)
  
  // rules := stringRules{
  //   minLen: 3,
  //   maxLen: 50,
  // }
  //
  // err := rules.validate(s)

  if err != nil {
    return "", fmt.Errorf("ActiveIntegrationId: %v", err)
  }

	return ActiveIntegrationId(s), nil
}

func NewAccountId(s string) (AccountId, error) {
  err := validators.String.IsNotBlank(s)

  if err != nil {
    return "", fmt.Errorf("AccountId: %v", err)
  }

	return AccountId(s), nil
}

func NewConfirmationTime(t time.Time) (ConfirmationTime, error) {
  err := validators.Time.IsNotZero(t)

  if err != nil {
    return ConfirmationTime{}, fmt.Errorf("ConfirmationTime: %v", err)
	}

	return ConfirmationTime(t), nil
}

func NewExtCompanyId(s string) (ExtCompanyId, error) {
  err := validators.String.IsNotBlank(s)

  if err != nil {
    return "", fmt.Errorf("ExtCompanyId: %v", err)
  }

	return ExtCompanyId(s), nil
}

func NewCompanyName(s string) (CompanyName, error) {
  err := validators.String.IsNotBlank(s)

  if err != nil {
    return "", fmt.Errorf("CompanyName: %v", err)
  }

	return CompanyName(s), nil
}

func (n CompanyName) ToString() string {
  return string(n)
}
