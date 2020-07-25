package config

import (
	"github.com/kelseyhightower/envconfig"
)

type configDef struct {
	AllowRegistration bool `split_words:"true" default:"false"`
	AllowAnonymous    bool `split_words:"true" default:"false"`

	Port int `default: "8080"`

	PostgresHost string `split_words:"true" default:"db"`
	PostgresPort int    `split_words:"true" default:"5432"`
	PostgresName string `split_words:"true" default:"dumbstored"`
	PostgresUser string `split_words:"true" default:"dumbstored"`
	PostgresPass string `split_words:"true" default:"dumbstored"`

	BcryptFactor int `split_words:"true" default:"12"`
}

var C configDef

func init() {
	envconfig.Process("dumbstored", &C)
}
