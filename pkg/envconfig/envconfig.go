package envconfig

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v11"
)

const appname = "SUPERADMIN"

func Load(spec any) {
	prefix := fmt.Sprintf("%s_", appname)

	err := env.ParseWithOptions(spec, env.Options{
		Prefix:                prefix,
		UseFieldNameByDefault: true,
		RequiredIfNoDef:       true,
	})

	if err != nil {
		log.Fatal(err.Error())
	}
}
