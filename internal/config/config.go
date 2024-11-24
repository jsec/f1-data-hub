package config

import "github.com/caarlos0/env/v11"

type Config struct {
	Host        string `env:"HOST" envDefault:"localhost"`
	Port        int    `env:"PORT" envDefault:"8080"`
	DatabaseURL string `env:"DATABASE_URL"`
}

func New() (Config, error) {
	var config Config
	if err := env.Parse(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}
