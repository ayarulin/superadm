package pg

import (
	. "superadmin.ru/accounts/core/interfaces"
	. "superadmin.ru/accounts/core/model"
	. "superadmin.ru/accounts/core/values"
)

var _ OrganizatoinRepository = &OrganizationsDb{}

type OrganizationsDb struct {
}

func NewOrganizationsDb(config interface{}) *OrganizationsDb {
	return &OrganizationsDb{}
}

func (d *OrganizationsDb) Delete(OrganizationId) error {
	return nil
}
func (d *OrganizationsDb) Save(Organization) error {
	return nil
}
