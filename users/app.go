package users

import (
	"context"

	"superadmin.ru/pkg/envconfig"
	"superadmin.ru/users/internal"
	"superadmin.ru/users/internal/core"
)

type app struct{ *ioc }

type ioc struct {
	config struct {
		JwtSecret         string
		YclientsUserToken string
	}

	usersRepo             core.UsersRepo
	jwtcoder              *application.JwtCoder
	yclientsSignValidator *application.YcSignValidator
}

type feature interface {
	call(*ioc, context.Context) error
}

func New() *app {
	ioc := ioc{}

	envconfig.Load(&ioc.config)

	ioc.jwtcoder = application.NewJwtCoder(ioc.config.JwtSecret)
	ioc.yclientsSignValidator = application.NewYcSignValidator(ioc.config.YclientsUserToken)

	return &app{&ioc}
}

func (a *app) Close() {
	panic("not implemented")
}

func (a *app) Run(ctx context.Context, feat feature) error {
	return feat.call(a.ioc, ctx)
}
