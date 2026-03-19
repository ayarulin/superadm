package model

import (
	"superadmin.ru/pkg/validation"
)

type YclientsRegistration struct {
	ID           int32
	ExtCompanyId string `validate:"required"`
	Name         string `validate:"required"`
	CompanyName  string `validate:"required"`
	Email        string `validate:"required,email"`
	PhoneNumber  string `validate:"required,e164"`
}

func NewYclientsRegistration(
	extCompanyId string,
	name string,
	companyName string,
	email string,
	phoneNumber string,
) (*YclientsRegistration, error) {

	model := YclientsRegistration{
		ExtCompanyId: extCompanyId,
		Name:         name,
		CompanyName:  companyName,
		Email:        email,
		PhoneNumber:  phoneNumber,
	}

	if err := validation.Struct(model); err != nil {
		return nil, err
	}

	return &model, nil
}
