package core

import (
	"superadmin.ru/pkg/clock"
	"time"
)

type ActiveIntegrationID int32
type ActiveIntegration struct {
	AccountId    AccountId
	CompanyName  string
	ExtCompanyId string
	ConfirmedAt  time.Time
}

type ActiveIntegrations interface {
	Get(ActiveIntegrationId) *ActiveIntegration
	Put(*ActiveIntegration) ActiveIntegrationId
	// Update(ActiveIntegrationId, ActiveIntegration)

	FilterByAccount(AccountId) []ActiveIntegrationId
}

// type ActiveIntegrationsFilter struct {
// 	AccountId    *AccountId
// 	CompanyName  *string
// 	ExtCompanyId string
// 	ConfirmedAt  time.Time
// }

func NewActiveIntegration(accId AccountId, extCompId ExtCompanyId, compName CompanyName) *ActiveIntegration {
	return &ActiveIntegration{
		AccountId:    accId,
		ExtCompanyId: extCompId.String(),
		CompanyName:  compName.String(),
		ConfirmedAt:  clock.Now(),
	}
}
