package configs

import "log"

func NewConfig(c ConfigLoader, envFilePath string) *Config {
	if c.ConsulAddress != "" {
		conf, err := c.loadFromConsul()
		if err != nil {
			log.Fatal("error loading config from consul", err)
		}
		return conf
	}

	return c.loadFromEnvFile(envFilePath)
}

func NewMockConfig() *Config {
	return &Config{
		RidesAPIKey: "test",
	}
}
