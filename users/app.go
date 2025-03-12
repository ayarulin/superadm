package users

import (
	"context"

	"superadmin.ru/pkg/envconfig"
	"superadmin.ru/users/service"

	"superadmin.ru/users/internal"
)

type app struct {
	config struct {
		YclientsUserToken string
		JwtSecret         string
	}
	commandIO      *core.CommandIO
	infrastructure struct{}
	services       struct {
		*service.JwtCoder
		*service.YcSignValidator
	}
}

func New() *app {
	app := app{}

	envconfig.Load(&app.config)

	// app.commandIO = &core.CommandIO{
	// 	UsersRepo:             nil,
	// 	YclientsSignValidator: ycsignvalidator.New(app.config.YclientsUserToken),
	// 	// JwtCoder:              jwtcoder.New(config.JwtSecret),
	// }

	app.services.JwtCoder = service.NewJwtCoder(app.config.JwtSecret)
	app.services.YcSignValidator = service.NewYcSignValidator(app.config.YclientsUserToken)

	return &app
}

func (a *app) Close() {
	panic("not implemented")
}

type executable interface {
	exec(*core.CommandIO, context.Context) error
}

func (a *app) Command(ctx context.Context, feat executable) error {
	return feat.exec(a.commandIO, ctx)
}
