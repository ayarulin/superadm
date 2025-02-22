package yclients

import (
	"superadmin.ru/yclients/infrastructure/http/yclientsapi"
	"superadmin.ru/yclients/infrastructure/postgres"
	"superadmin.ru/yclients/internal/command"
)

type yclientsApp struct {
	db *postgres.PostgresDB
	*command.Command
}

func NewApp() *yclientsApp {
  config := LoadConfig()

	postgresDB := postgres.Open(
		config.PostgresHost,
		config.PostgresDb,
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresPort,
	)

	yclientsapi := yclientsapi.New(config.APIKey)

	cmd := command.NewCommand(
		postgresDB.ActiveIntegrations,
		yclientsapi,
	)

	return &yclientsApp{
		postgresDB,
		cmd,
	}
}

func (y *yclientsApp) Close() {
	y.db.Close()
}
