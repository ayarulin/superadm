package model

import (
	. "superadmin.ru/modules/accounts/internal/core/values"
	"superadmin.ru/pkg/id"
)

type Organization struct {
	Id      OrganizationId
	Name    OrganizationName
	OwnerId UserId
}

func NewOrganization(ownerId UserId, name OrganizationName) Organization {
	return Organization{
		Id:      OrganizationId(id.Random()),
		OwnerId: ownerId,
		Name:    name,
	}
}
