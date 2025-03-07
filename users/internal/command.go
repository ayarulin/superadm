package internal

import(
  "context"
)


type YclientsIntegrationActivator interface {
  Call(ctx context.Context, userId int32, ycCompanyId string) error
}

type Command struct {
  yclientsIntegrationActivator YclientsIntegrationActivator
}


func NewCommand() *Command {
  return &Command{}
}
