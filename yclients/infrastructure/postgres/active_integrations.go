package postgres

import (
	"context"
	"log"

	. "superadmin.ru/yclients/infrastructure/postgres/dao"
	"superadmin.ru/yclients/internal/core"
)

type activeIntegrations struct {
	*Queries
}

// FilterByAccount implements core.ActiveIntegrations.
func (a *activeIntegrations) FilterByAccount(core.AccountId) []core.ActiveIntegrationId {
	panic("unimplemented")
}

// Create implements core.ActiveIntegrations.
func (a *activeIntegrations) Put(ai *core.ActiveIntegration) core.ActiveIntegrationId {
	id, err := a.Queries.CreateActiveIntegration(context.TODO(), CreateActiveIntegrationParams{
		AccountID:        ai.AccountId.Int32(),
		ExtCompanyID:     ai.ExtCompanyId,
		CompanyName:      ai.CompanyName,
		ConfirmationTime: ai.ConfirmedAt,
	})

	if err != nil {
		log.Panic(err)
	}

	return core.NewActiveIntegrationId(id)
}

// Get implements core.ActiveIntegrations.
func (a *activeIntegrations) Get(id core.ActiveIntegrationId) *core.ActiveIntegration {
	rec, err := a.Queries.GetActiveIntegration(context.TODO(), id.Int32())

	if err != nil {
		log.Panic(err)
	}

	return &core.ActiveIntegration{
		AccountId:    core.AccountId(rec.AccountID),
		ExtCompanyId: rec.ExtCompanyID,
		CompanyName:  rec.CompanyName,
		ConfirmedAt:  rec.ConfirmationTime,
	}
}

