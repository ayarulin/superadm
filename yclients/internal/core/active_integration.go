package core

import (
	"time"
)

type ActiveIntegrationId string
type ActivationTime time.Time
type ConfirmationTime time.Time
type SyncTime time.Time
type AccountId string
type ExtCompanyId string
type CompanyName string

type ActiveIntegration struct {
	id             ActiveIntegrationId
	accountId      AccountId
	extCompanyId   ExtCompanyId
	companyName    CompanyName
	activationTime ActivationTime
	syncTime       SyncTime
}

func NewActiveIntegrationId(val string) (ActiveIntegrationId, error) {
	return ActiveIntegrationId(val), nil
}

func NewActivationTime(now, val time.Time) (ActivationTime, error) {
  time := ActivationTime(val)

  if (now < val) {
    return time, fmt.
  }
}
