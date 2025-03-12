package yclients

import "superadmin.ru/pkg/envconfig"

type config struct {
	YclientsDatabaseUrl string
	YclientsUserToken   string
}

func loadConfig() config {
	config := config{}
	envconfig.Load(&config)

	return config
}
