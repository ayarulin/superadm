package core

import "context"

type YclientsService interface {
	ConfirmRegistration(context.Context, ExtCompanyId) error
	GetCompany(context.Context, ExtCompanyId) (CompanyName, error)
}

