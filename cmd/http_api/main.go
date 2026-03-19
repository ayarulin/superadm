package main

import (
	"fmt"
	"net/http"

	"superadmin.ru/users"
	"superadmin.ru/internal/eventbus"
)

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {

	// TODO: move up to main()

	eventBus := eventbus.EventBus{}
	usersDomain := users.New(&eventBus)


	users.New(&eventBus)

	_, err := usersDomain.YclientsRegistrationAccept.Call(
		req.Context(),
		users.YclientsRegistrationAcceptInput{
			SalonId:      "",
			SalonName:    "",
			UserName:     "",
			UserEmail:    "",
			UserPhone:    "",
			UserData:     "",
			UserDataSign: "",
		},
	)

	fmt.Fprintf(w, "result: %v \n", err)
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
