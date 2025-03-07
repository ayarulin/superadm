package users

import (
	"superadmin.ru/users/internal"
)

type app struct {
	*internal.Command
}

func NewApp() *app {
  _ = LoadConfig()

	cmd := internal.NewCommand()

	return &app{
    cmd,
  }
}

func (a *app) Close() {
	// a.db.Close()
}
