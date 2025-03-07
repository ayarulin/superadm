package users

import (
	"context"
	"superadmin.ru/users/internal"
)

type app struct {
	// db
  command
}

func New() *app {
	_ = LoadConfig()

	return &app{
	 command{},
	}
}

func (a *app) Close() {
	panic("not implemented")
}

type command struct {
	usersRepo core.UsersRepo
}

type handlerFunc func(context.Context) (any, error)
type executable interface {
	Handler(*command) handlerFunc
}

func (c *command) Execute(ctx context.Context, e executable) (any, error){
  return e.Handler(c)(ctx)
}


type Command[T, R any] interface {
	Handle(T) (R, error)
}
