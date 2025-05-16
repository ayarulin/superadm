package users

import (
	"context"
	"fmt"

	"superadmin.ru/app/users/events"
	"superadmin.ru/app/users/internal"
	"superadmin.ru/app/users/internal/core"
	"superadmin.ru/app/users/internal/repository"
	"superadmin.ru/infrastructure/eventbus"
	"superadmin.ru/infrastructure/postgres"
)

type YclientsRegistrationAcceptHandler struct {
	signValidator         *internal.YcSignValidator
	txManager             *postgres.TxManager
	yclientsRegistrations *repository.YclientsRegistrations
	eventBus              *eventbus.EventBus
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
	signValid := h.signValidator.ValidateSign(input.UserData, input.UserDataSign)

	if !signValid {
		return fmt.Errorf("Signature invalid")
	}

	reg, err := core.NewYclientsRegistration(
		input.SalonID,
		input.UserName,
		input.SalonName,
		input.UserEmail,
		input.UserPhone,
	)

	if err != nil {
		return err
	}

	err = h.txManager.WithTx(ctx,
		func(ctx context.Context) error {
			err := h.yclientsRegistrations.Save(ctx, reg)

			if err != nil {
				return err
			}

			err = h.eventBus.Publish(
				events.YclientsRegistrationAccepted{
					YclientsRegstrationId: reg.ID,
					SalonId:               reg.ExtCompanyId,
				})

			return err
		})

	return err
}
