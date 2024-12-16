package accounts

import (
	. "superadmin.ru/modules/accounts/internal/core"
	. "superadmin.ru/modules/accounts/internal/core/values"
	"superadmin.ru/modules/accounts/internal/infrastructure/pg"
)

type AccountsApp struct {
	core *AccountsCore
}

func NewAccountsApp(config interface{}) *AccountsApp {

	orgRepo := pg.NewOrganizationsDb(config)
	core := NewAccountsCore(orgRepo)

	return &AccountsApp{
		core: core,
	}
}

func (a *AccountsApp) AddOrganization(name string, ownerId string) (string, error) {
	orgName, err := NewOrganizationName(name)

	if err != nil {
		return "", err
	}

	userId, err := NewUserId(ownerId)

	if err != nil {
		return "", err
	}

	org, err := a.core.AddOrganization(orgName, userId)

	if err != nil {
		return "", err
	}

	return string(org.Id), nil
}
