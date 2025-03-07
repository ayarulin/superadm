package yclientsapi

import (
	"context"

	"errors"
	"superadmin.ru/yclients/internal/core"
)

type YClientsAPI struct {
	apiKey string
}

func New(apiKey string) *YClientsAPI {
	return &YClientsAPI{
		apiKey,
	}
}

// ConfirmRegistration implements core.YclientsService.
func (y *YClientsAPI) ConfirmRegistration(context.Context, core.ExtCompanyId) error {
	return nil
}

// GetCompany implements core.YclientsService.
func (y *YClientsAPI) GetCompany(context.Context, core.ExtCompanyId) (core.CompanyName, error) {
	return core.CompanyName(""), errors.New("service error")
}
