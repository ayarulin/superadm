package main

import (
	"context"
	"fmt"
	"net/http"

	"superadmin.ru/pkg/envconfig"
	"superadmin.ru/yclients"
)

func main() {
	cfg := &envconfig.Config

	ycliensApp := yclients.New(
		yclients.Config{
			PostgresDb:       cfg.PostgresDB,
			PostgresUser:     cfg.PostgresUser,
			PostgresPassword: cfg.PostgresPassword,
			YclientsAPIKey:   cfg.YclientsAPIKey,
		},
	)

  ctx := context.Background()

  _, _ = ycliensApp.ActivateIntegration(ctx, "accid", "extCompId")

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
