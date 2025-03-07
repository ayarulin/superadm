package main

import (
	"fmt"
	"net/http"

	"superadmin.ru/yclients"
)



func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	ycliensApp := yclients.NewApp()

  _, _ = ycliensApp.ActivateIntegration(req.Context(), 1, "extCompId")

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
