package accounts

import (
	"regexp"
	"testing"

	"superadmin.ru/accounts/core/internal"
)

type InmemoryOrganizationRepo struct {
	items map[internal.OrganizationId]internal.Organization
}

func NewInmemoryOrganizationRepo() InmemoryOrganizationRepo {
	return InmemoryOrganizationRepo{
		items: make(map[internal.OrganizationId]internal.Organization),
	}
}

func (r *InmemoryOrganizationRepo) Delete(internal.OrganizationId) error {
	return nil
}

func (r *InmemoryOrganizationRepo) Save(*internal.Organization) error {
	return nil
}

func TestAddOrganization(t *testing.T) {
	repo := NewInmemoryOrganizationRepo()
	accounts := NewAccountsCore(&repo)

	accounts.AddOrganization(name, owner_id)

	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}
