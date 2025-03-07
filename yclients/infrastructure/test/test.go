package test

import (
	"context"
	"errors"

	. "superadmin.ru/yclients/internal/core"
)

type fakeApi struct{
  data map[ExtCompanyId]CompanyName
}

func NewYclientsFakeApi() *fakeApi {
	return &fakeApi{
    data: map[ExtCompanyId]CompanyName{},
  }
}

func (f *fakeApi) ConfirmRegistration(ctx context.Context, id ExtCompanyId) error {
  return nil
}

func (f *fakeApi) GetCompany(ctx context.Context, id ExtCompanyId) (CompanyName, error) {
  name, ok := f.data[id] 

  if !ok {
    return CompanyName(""), errors.New("not found")
  }

  return name, nil
}

func (f *fakeApi) AddCompany(id, name string) {
  f.data[ExtCompanyId(id)] = CompanyName(name)
}

