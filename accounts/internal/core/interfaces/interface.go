package interfaces

import (
	. "superadmin.ru/accounts/internal/core/model"
	. "superadmin.ru/accounts/internal/core/values"
)

type OrganizatoinRepository interface {
	Save(Organization) error
	Delete(OrganizationId) error
}

type AuthServiceApi interface {
}
