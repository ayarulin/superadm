package envconfig

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

const appname = "SUPERADMIN"

func Load(scope string, spec interface{}) {
	loadEnvFile()

	scope = strings.ToUpper(scope)
	prefix := fmt.Sprintf("%s_%s_", appname, scope)

	err := env.ParseWithOptions(spec, env.Options{
		Prefix:                prefix,
		UseFieldNameByDefault: true,
		RequiredIfNoDef:       true,
	})

	if err != nil {
		log.Fatal(err.Error())
	}
}

func loadEnvFile() {
	env := os.Getenv("SUPERADMIN_ENV")

	if env != "test" && env != "dev" {
		return
	}

	err := godotenv.Load("../env/" + env + ".env")

	if err != nil {
		log.Fatal(err.Error())
	}
}
