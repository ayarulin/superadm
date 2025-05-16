package users

import (
	"superadmin.ru/app/users/internal"
	"superadmin.ru/app/users/internal/repository"
	"superadmin.ru/infrastructure/eventbus"
	"superadmin.ru/infrastructure/postgres"
	"superadmin.ru/pkg/envconfig"
)

type Users struct {
	*YclientsRegistrationAcceptHandler
}

type config struct {
	UsersDatabaseUrl  string
	JwtSecret         string
	YclientsUserToken string
}

func New(eventBus *eventbus.EventBus) *Users {
	config := config{}
	envconfig.Load(&config)

	db := postgres.Open(config.UsersDatabaseUrl)

	return &Users{
		&YclientsRegistrationAcceptHandler{
			eventBus:              eventBus,
			signValidator:         internal.NewYcSignValidator(config.YclientsUserToken),
			txManager:             postgres.NewTxManager(db),
			yclientsRegistrations: repository.NewYclientsRegistrations(db),
		},
	}

}
