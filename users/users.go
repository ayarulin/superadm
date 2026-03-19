package users

import (
	"superadmin.ru/internal/envconfig"
	"superadmin.ru/internal/postgres"
	"superadmin.ru/users/commands"
	"superadmin.ru/users/crypto"
	"superadmin.ru/users/repository"

	yclientsAPI "superadmin.ru/yclients/api"
)

type Users struct {
	*commands.YclientsRegistrationAcceptHandler
}

type config struct {
	UsersDatabaseUrl  string
	JwtSecret         string
	YclientsUserToken string
}

func New(yclientsService yclientsAPI.API) *Users {
	config := config{}
	envconfig.Load(&config)

	db := postgres.Open(config.UsersDatabaseUrl)

	return &Users{
		&commands.YclientsRegistrationAcceptHandler{
			SignValidator:         crypto.NewYcSignValidator(config.YclientsUserToken),
			TxManager:             postgres.NewTxManager(db),
			YclientsRegistrations: repository.NewYclientsRegistrations(db),
			YclientsService:       yclientsService,
		},
	}
}
