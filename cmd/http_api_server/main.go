package main

import (
	"fmt"
	"net/http"

	"superadmin.ru/users"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	usersApp := users.New()
	defer usersApp.Close()

	err := usersApp.Run(req.Context(),
		&users.YclientsRegistrationInitCmd{
			ExtCompanyId: "",
			Name:         "",
			CompanyName:  "",
			Email:        "",
			PhoneNumber:  "",
			UnsignedJSON: "",
			Sign:         "",
		})

	fmt.Fprintf(w, "result: %v \n", err)
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
