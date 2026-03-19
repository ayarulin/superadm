package users

import (
	"superadmin.ru/internal/envconfig"
	"superadmin.ru/internal/postgres"
	"superadmin.ru/users/commands"
	"superadmin.ru/users/crypto"
	"superadmin.ru/users/repository"
)

type API struct {
	db     *postgres.DB
	config struct {
		UsersDatabaseUrl  string
		JwtSecret         string
		YclientsUserToken string
	}
}

func New() *API {
	api := API{}

	envconfig.Load(&api.config)

	api.db = postgres.Open(api.config.UsersDatabaseUrl)

	return &api
}
