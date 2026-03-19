package internal

import (
  "context"
	. "superadmin.ru/yclients/internal/core"
)

type Command struct {
  activeIntegrations ActiveIntegrations
  yclientsService YclientsService
}

type YclientsService interface {
	ConfirmRegistration(ctx context.Context, companyId ExtCompanyId) error
	GetCompany(ctx context.Context, companyId ExtCompanyId) (CompanyName, error)
}

func NewCommand(
	ais ActiveIntegrations,
	ys YclientsService,
) *Command {
  return &Command{
    activeIntegrations: ais,
    yclientsService: ys,
  }
}
