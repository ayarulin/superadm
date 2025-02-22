package command

import (
	"context"
	"fmt"
	"log"

	"superadmin.ru/pkg/clock"
	"superadmin.ru/pkg/id"
	. "superadmin.ru/yclients/internal/command/core"
)

func (c *Command) ActivateIntegration(
	ctx context.Context,
	accountId string,
	extCompanyId string,
) (intId string, err error) {
	{
		accountId, err := NewAccountId(accountId)

		if err != nil {
			return "", validationError(err)
		}

		extCompanyId, err := NewExtCompanyId(extCompanyId)
		if err != nil {
			return "", validationError(err)
		}

		confTime, err := NewConfirmationTime(clock.Now())

		if err != nil {
      log.Fatal(err.Error())
		}

		err = c.yclientsService.ConfirmRegistration(ctx, extCompanyId)

		if err != nil {
			return "", err
		}

		compName, err := c.yclientsService.GetCompany(ctx, extCompanyId)

		if err != nil {
			return "", err
		}

    intId = id.Random()
    id, err := NewActiveIntegrationID(intId)

		if err != nil {
      log.Fatal(err.Error())
		}

		actInt := NewActiveIntegration(
      id,
			accountId,
			extCompanyId,
			compName,
			confTime,
		)

    err = c.activeIntegrations.Create(actInt)

		if err != nil {
      log.Fatal(err.Error())
		}

		return intId, nil
	}
}

func validationError(err error) error {
	return fmt.Errorf("Validation error: %v", err)
}
