package internal

import (
	"context"
	"fmt"
  "log"

	. "superadmin.ru/yclients/internal/core"
)

func (c *Command) ActivateIntegration(
	ctx context.Context,
	accId int32,
	extCompanyId string,
) (id int32, err error) {
	{
		extCompanyId, err := NewExtCompanyId(extCompanyId)

		if err != nil {
			return 0, validationError(err)
		}

		err = c.yclientsService.ConfirmRegistration(ctx, extCompanyId)

		if err != nil {
      // if already confirmed try to find in db or create new one
			return 0, err
		}

		compName, err := c.yclientsService.GetCompany(ctx, extCompanyId)

		if err != nil {
      log.Printf("Failed to get company name: %v", err)

      compName = CompanyName(fmt.Sprintf("Salon #%s", extCompanyId.String()))
		}

		actInt := NewActiveIntegration(
			NewAccountId(accId),
			extCompanyId,
			compName,
		)

		actIntId := c.activeIntegrations.Put(actInt)

		return actIntId.Int32(), nil
	}
}

func validationError(err error) error {
	return fmt.Errorf("Validation error: %v", err)
}
