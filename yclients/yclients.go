package yclients

import (
	"superadmin.ru/yclients/infrastructure/http/yclientsapi"
	"superadmin.ru/yclients/infrastructure/postgres"
	"superadmin.ru/yclients/internal"
)

type app struct {
	db *postgres.PostgresDB
	*internal.Command
}

func NewApp() * app{
  config := LoadConfig()

	postgresDB := postgres.Open(config.DatabaseUrl)
	yclientsapi := yclientsapi.New(config.APIKey)

	cmd := internal.NewCommand(
		postgresDB.ActiveIntegrations,
		yclientsapi,
	)

	return &app{
		postgresDB,
		cmd,
	}
}


func (y *app) Close() {
	y.db.Close()
}
