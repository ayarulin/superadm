package users

import (
	"context"
	"fmt"
	"superadmin.ru/users/internal/core"
)

type YclientsRegistrationInitCmd struct {
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

func (cmd *YclientsRegistrationInitCmd) call(ioc *ioc, ctx context.Context) error {
	if !ioc.yclientsSignValidator.ValidateSign(cmd.UnsignedJSON, cmd.Sign) {
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

	token, err := ioc.jwtcoder.EncodeYclientsRegistration(reg)

	if err != nil {
		return err
	}

	cmd.result.Token = token
	return nil
}

func (cmd *YclientsRegistrationInitCmd) Result() string {
	return cmd.result.Token
}
