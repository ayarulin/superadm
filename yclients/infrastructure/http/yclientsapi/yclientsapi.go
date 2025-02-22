package yclientsapi

import (
	"context"
	"superadmin.ru/yclients/internal/command/core"
)

type YClientsAPI struct{}


func New(apiKey string) *YClientsAPI {
	return &YClientsAPI{}
}

// ConfirmRegistration implements core.YclientsService.
func (y *YClientsAPI) ConfirmRegistration(context.Context, core.ExtCompanyId) error {
	panic("unimplemented")
}

// GetCompany implements core.YclientsService.
func (y *YClientsAPI) GetCompany(context.Context, core.ExtCompanyId) (core.CompanyName, error) {
	panic("unimplemented")
}
