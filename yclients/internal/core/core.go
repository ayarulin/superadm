package core

import (
	"context"
	"time"
  "superadmin.ru/internal/id"
)

const (
  // ErrYclientsServiceUnavailabe = errors.Err
)

type YClientsCore struct { 
  repository Repository
  yClientsGateway YClientsGateway
}

// type Repository interface {
//   SaveActiveIntegration(c context.Context, integration *ActiveIntegration) error
// }

type YClientsGateway interface {
  ConfirmRegistration(context.Context,  ExtCompanyId) error 
  GetCompany(context.Context, ExtCompanyId) (CompanyName, error)
}

func (c *YClientsCore) ActivateIntegration(ctx context.Context, accountId AccountId, extCompanyId ExtCompanyId) (ActiveIntegration, error) {
  err := c.yClientsGateway.ConfirmRegistration(ctx, extCompanyId)

  if err != nil  {
    return ActiveIntegration{}, err
  }

  companyName, err := c.yClientsGateway.GetCompany(ctx, extCompanyId)

  if err != nil {
    return ActiveIntegration{}, err
  }

  activeInt := ActiveIntegration{
		id:             ActiveIntegrationId(id.Random()),
		accountId:      accountId,
		extCompanyId:   extCompanyId,
		companyName:    companyName,
		activationTime: ActivationTime(time.Now()),
	}

  err = c.repository.SaveActiveIntegration(ctx, &activeInt)

  return activeInt, err
}

type Event func(i *ActiveIntegration)

func ActivateIntegrationCmd(accountId AccountId, extCompanyId ExtCompanyId) Event {
  return func(i *ActiveIntegration) {
    i.accountId = accountId
    i.extCompanyId = extCompanyId
  }
}


