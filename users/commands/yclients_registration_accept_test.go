package commands_test

import (
	"context"
	"os"
	"testing"

	"superadmin.ru/internal/shared/crypto"
	"superadmin.ru/users"
	"superadmin.ru/users/commands"
)

func TestYclientsRegistrationInitCmd_SignCheck(t *testing.T) {
	token := os.Getenv("SUPERADMIN_YCLIENTS_USER_TOKEN")

	usersDomain := users.New(&eventbus.EventBus{})

	err := usersDomain.YclientsRegistrationAcceptHandler.Call(
		context.Background(),

		commands.YclientsRegistrationAcceptInput{
			SalonID:      "Salon#1",
			SalonName:    "Salon",
			UserName:     "User#1",
			UserEmail:    "user1@example.com",
			UserPhone:    "+71234567890",
			UserData:     "{some: 'value'}",
			UserDataSign: "",
		})

	if err == nil || err.Error() != "Signature invalid" {
		t.Fatalf("Signature check should not be passed but it does")
	}

	err = usersDomain.YclientsRegistrationAcceptHandler.Call(
		context.Background(),
		users.YclientsRegistrationAcceptInput{
			SalonID:      "Salon#1",
			SalonName:    "Salon",
			UserName:     "User#1",
			UserEmail:    "user1@example.com",
			UserPhone:    "+71234567890",
			UserData:     "{some: 'value'}",
			UserDataSign: crypto.Sign("{some: 'value'}", token),
		})

	if err != nil {
		t.Fatalf(`%v`, err)
	}
}
