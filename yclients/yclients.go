package yclientsapp

import (
	"context"

	"superadmin.ru/yclients/internal/core"
)

type YClientsApp struct {
  core core.YClientsCore
}

func NewYclients() *YClientsApp {
	return &YClientsApp{
    core: core.YClientsCore{},
	}
}

func (app *YClientsApp) ActivateIntegration(c context.Context, accountId string, extCompanyId string) string {

  accId := core.AccountId(accountId)
  extCompId := core.ExtCompanyId(extCompanyId)

  event := core.ActivateIntegrationCmd(accId, extCompId)

	return "id"
}
