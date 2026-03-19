package commands

import (
	"context"
	"fmt"
	"superadmin.ru/internal/postgres"
	"superadmin.ru/users/crypto"
	"superadmin.ru/users/model"
	"superadmin.ru/users/repository"

	yclients "superadmin.ru/yclients/api"
)

type YclientsRegistrationAcceptHandler struct {
	SignValidator         *crypto.YcSignValidator
	TxManager             *postgres.TxManager
	YclientsRegistrations *repository.YclientsRegistrations
	Yclients              *yclients.API
}

type YclientsRegistrationAcceptInput struct {
	SalonID      string
	SalonName    string
	UserName     string
	UserEmail    string
	UserPhone    string
	UserData     string
	UserDataSign string
}

func (h *YclientsRegistrationAcceptHandler) Call(ctx context.Context, input YclientsRegistrationAcceptInput) error {
	signValid := h.SignValidator.ValidateSign(input.UserData, input.UserDataSign)

	if !signValid {
		return fmt.Errorf("Signature invalid")
	}

	reg, err := model.NewYclientsRegistration(
		input.SalonID,
		input.UserName,
		input.SalonName,
		input.UserEmail,
		input.UserPhone,
	)

	if err != nil {
		return err
	}

	err = h.TxManager.WithTx(ctx,
		func(ctx context.Context) error {
			err := h.YclientsRegistrations.Save(ctx, reg)

			if err != nil {
				return err
			}

			err = h.Yclients.ActivateIntegration(reg.ID, reg.ExtCompanyId)

			return err
		})

	return err
}
