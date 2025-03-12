package users

import (
	"context"
	"fmt"
	"superadmin.ru/users/internal"
)

type InitYclientsRegistration struct {
	ExtCompanyId string
	Name         string
	CompanyName  string
	Email        string
	PhoneNumber  string
	UnsignedJSON string
	Sign         string
result struct {
		Token string
	}
}

func (cmd *InitYclientsRegistration) exec(app *app, ctx context.Context) error {
	if !app.services.YcSignValidator.ValidateSign(cmd.UnsignedJSON, cmd.Sign) {
		return fmt.Errorf("Signature invalid")
	}

	reg, err := core.NewYclientsRegistration(
		cmd.ExtCompanyId,
		cmd.Name,
		cmd.CompanyName,
		cmd.Email,
		cmd.PhoneNumber,
	)

	if err != nil {
		return err
	}

  token, err := app.services.JwtCoder.EncodeYclientsRegistration(reg)

  if err != nil {
		return err
  }

	cmd.result.Token = token
	return nil
}

func (cmd *InitYclientsRegistration) Result() string {
	return cmd.result.Token
}
func (cmd *InitYclientsRegistration) exec(io *core.CommandIO, ctx context.Context) error {
	if !io.YclientsSignValidator.ValidateSign(cmd.UnsignedJSON, cmd.Sign) {
		return fmt.Errorf("Signature invalid")
	}

	reg, err := core.NewYclientsRegistration(
		cmd.ExtCompanyId,
		cmd.Name,
		cmd.CompanyName,
		cmd.Email,
		cmd.PhoneNumber,
	)

	if err != nil {
		return err
	}

  token, err := io.JwtCoder.EncodeYclientsRegistration(reg)

  if err != nil {
		return err
  }

	cmd.result.Token = token
	return nil
}

func (cmd *InitYclientsRegistration) Result() string {
	return cmd.result.Token
}
