package api

import (
	"superadmin.ru/internal/envconfig"
	"superadmin.ru/yclients/infrastructure/postgres"
)

type API struct {
	db     *postgres.PostgresDB
	config struct {
		YclientsDatabaseUrl string
		YclientsUserToken   string
	}
}

func New() *API {
	api := API{}
	envconfig.Load(&api.config)
	api.db = postgres.Open(api.config.YclientsDatabaseUrl)

	return &api
}

func (a *API) ActivateIntegration(registrationID int32, salonID string) error {
	// call command
	return nil
}
