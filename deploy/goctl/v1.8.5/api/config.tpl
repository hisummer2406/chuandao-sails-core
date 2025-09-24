package config

import (
    "chuandao-sails-core/common/sls"
    {{.authImport}}
)

type Config struct {
	rest.RestConf
	SLS sls.Config
	{{.auth}}
	{{.jwtTrans}}
}
