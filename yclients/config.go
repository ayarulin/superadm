package yclients

import "superadmin.ru/pkg/envconfig"

type config struct {
	DatabaseUrl string
	APIKey      string
}

func LoadConfig() config {
	config := config{}
	envconfig.Load("YCLIENTS", &config)

	return config
}
