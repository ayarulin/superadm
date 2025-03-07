package users_test

import (
	"context"
	"testing"

	"superadmin.ru/users"
)

func TestInitYclientsRegistration(t *testing.T) {
	usersApp := users.New()

	token, err := usersApp.Execute(
		context.Background(),
		users.InitYclientsRegistration{
			ExtCompanyId: "",
			Name:         "",
			CompanyName:  "",
			Email:        "",
			PhoneNumber:  "",
			Signature:    "",
		},
	)

	{
		var token string = token.(string)

		if err != nil {
			t.Fatalf(`%v`, err)
		}

		if token == "" {
			t.Fatalf(`Token is blank: %s`, token)
		}
	}
}
