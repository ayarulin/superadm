package accounts

import (
	. "superadmin.ru/modules/accounts/internal/core/model"
	. "superadmin.ru/modules/accounts/internal/core/values"
)

func (s *AccountsCore) AddOrganization(name OrganizationName, owner_id UserId) (Organization, error) {
	org := Organization{
		Name:    name,
		OwnerId: owner_id,
	}

	err := s.organizationsRepo.Save(org)

	return org, err
}
