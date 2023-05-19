package config

import (
	"github.com/caarlos0/env/v8"
	"github.com/disgoorg/log"
	"github.com/joho/godotenv"
)

type Config struct {
	TLS TLS
}

type TLS struct {
	CertFile string `env:"TLS_CERT_FILE,required"`
	KeyFile  string `env:"TLS_KEY_FILE,required"`
}

func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}
