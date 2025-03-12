package users_test

import (
	"context"
	"os"
	"testing"

	"superadmin.ru/pkg/crypto"
	. "superadmin.ru/users"
)

func TestInitYclientsRegistrationSignCheck(t *testing.T) {
	token := os.Getenv("SUPERADMIN_YCLIENTS_USER_TOKEN")

	feat := InitYclientsRegistration{
		ExtCompanyId: "comp#1",
		Name:         "username#1",
		CompanyName:  "compname#1",
		Email:        "user#1@example.com",
		PhoneNumber:  "+71234567890",
		UnsignedJSON: "{some: value}",
		Sign:         "",
	}


  usersApp := New()

  err := usersApp.Command(context.Background(), &feat)

	if err == nil || err.Error() != "Signature invalid" {
		t.Fatalf("Signature check should not be passed but it does")
	}

	feat.Sign = crypto.Sign(feat.UnsignedJSON, token)

	err = usersApp.Command(context.Background(), &feat)

	if err != nil {
		t.Fatalf(`%v`, err)
	}

	if feat.Result() == "" {
		t.Fatalf(`Result token is blank: %s`, token)
	}
}
