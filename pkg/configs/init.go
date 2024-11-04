package configs

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

func NewConfig(envFilePath string) *Config {
	if envFilePath != "" {
		err := godotenv.Load(envFilePath)
		if err != nil {
			log.Println("error loading config from file", err)
		}
	}

	cfg := new(Config)
	if err := env.Parse(cfg); err != nil {
		log.Fatal("error parsing config from file", err)
	}

	return cfg
}

func NewMockConfig() *Config {
	return &Config{}
}
