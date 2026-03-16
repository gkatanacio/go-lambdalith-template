package sample

import "github.com/kelseyhightower/envconfig"

type Config struct {
	HelloWho string `envconfig:"HELLO_WHO" required:"true"`
}

func ConfigFromEnv() Config {
	var config Config
	envconfig.MustProcess("", &config)

	return config
}
