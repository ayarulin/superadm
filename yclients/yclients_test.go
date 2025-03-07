package yclients_test

import (
	"context"
	"testing"

	"superadmin.ru/yclients"
)

func TestActivateIntegration(t *testing.T) {
	yclientsApp := yclients.NewApp()

	ctx := context.Background()

	intId, err := yclientsApp.Command.ActivateIntegration(ctx, 1, "extCompany#1")

  if intId == 0 || err != nil {
      t.Fatalf(`%v`, err)
  }
}
