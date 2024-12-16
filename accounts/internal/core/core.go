package accounts

import . "superadmin.ru/accounts/internal/core/interfaces"

type AccountsCore struct {
	organizationsRepo OrganizatoinRepository
}

func NewAccountsCore(organizationRepository OrganizatoinRepository) *AccountsCore {
	return &AccountsCore{
		organizationsRepo: organizationRepository,
	}
}
