package yclients

import "superadmin.ru/pkg/envconfig"

type config struct {
	PostgresDb       string
	PostgresUser     string
	PostgresPassword string
	PostgresHost     string
	PostgresPort     int
	APIKey           string
}

func LoadConfig() config {
  config := config{}
	envconfig.Load("YCLIENTS", &config)

  return config
}
