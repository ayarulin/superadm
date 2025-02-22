package command

import (
	. "superadmin.ru/yclients/internal/command/core"
)

type Command struct {
  activeIntegrations ActiveIntegrations
  yclientsService YclientsService
}

func NewCommand(
	ais ActiveIntegrations,
	ys YclientsService,
) *Command {
  return &Command{
    activeIntegrations: ais,
    yclientsService: ys,
  }
}
