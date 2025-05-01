package users_test

import (
	"context"
	"os"
	"testing"

	"superadmin.ru/pkg/crypto"
	. "superadmin.ru/users"
)

func TestYclientsRegistrationInitCmd_SignCheck(t *testing.T) {
	token := os.Getenv("SUPERADMIN_YCLIENTS_USER_TOKEN")
	usersApp := New()
  
	cmd := YclientsRegistrationInitCmd{
		ExtCompanyId: "comp#1",
		Name:         "username#1",
		CompanyName:  "compname#1",
		Email:        "user#1@example.com",
		PhoneNumber:  "+71234567890",
		UnsignedJSON: "{some: value}",
		Sign:         "",
	}

	err := usersApp.Run(context.Background(), &cmd)

	if err == nil || err.Error() != "Signature invalid" {
		t.Fatalf("Signature check should not be passed but it does")
	}

	cmd.Sign = crypto.Sign(cmd.UnsignedJSON, token)

	err = usersApp.Run(context.Background(), &cmd)

	if err != nil {
		t.Fatalf(`%v`, err)
	}

	if cmd.Result() == "" {
		t.Fatalf(`Result token is blank: %s`, token)
	}
}
