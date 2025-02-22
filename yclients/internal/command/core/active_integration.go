package core

import (
	"fmt"
  "superadmin.ru/pkg/validators"
)

type ActiveIntegration struct {
	id               IntegrationID
	accountId        AccountId
	extCompanyId     ExtCompanyId
	companyName      CompanyName
	confirmationTime ConfirmationTime
}

type ActiveIntegrations interface {
  Get(ActiveIntegrationId) (*ActiveIntegration, error)
	Create(*ActiveIntegration) error
	// Update(*ActiveIntegration) error
	// Delete(ActiveIntegrationId) error
}

// type ActiveIntegrationsFinder interface {
// 	Find(ActiveIntegrationId) (*ActiveIntegrationQueryModel, error)
// 	FindByAccount(AccountId) []ActiveIntegration
// }

type IntegrationID string

func NewActiveIntegration(
	id IntegrationID,
	accId AccountId,
	extCompId ExtCompanyId,
	compName CompanyName,
	confirmationTime ConfirmationTime,
) *ActiveIntegration {
	return &ActiveIntegration{
		id:               id,
		accountId:        accId,
		extCompanyId:     extCompId,
		companyName:      compName,
		confirmationTime: confirmationTime,
	}
}

func NewActiveIntegrationID(s string) (IntegrationID, error) {
  err := validators.String.IsNotBlank(s)

  if err != nil {
    return "", fmt.Errorf("ActiveIntegrationId: %v", err)
  }

	return IntegrationID(s), nil
}
