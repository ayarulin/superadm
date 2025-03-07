package users

import "superadmin.ru/pkg/envconfig"

type config struct {
}

func LoadConfig() config {
	config := config{}
	envconfig.Load("USERS", &config)

	return config
}
