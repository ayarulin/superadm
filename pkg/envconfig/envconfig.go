package envconfig

import (
	"fmt"
	"log"
	"strings"

	"github.com/caarlos0/env/v11"
)

const appname = "SUPERADMIN"

func Load(scope string, spec interface{}) {

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
