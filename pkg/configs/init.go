package configs

import "log"

func NewConfig(c ConfigLoader) *Config {
	if c.ConsulAddress != "" {
		conf, err := c.loadFromConsul()
		if err != nil {
			log.Fatal("error loading config from consul", err)
		}
		return conf
	}

	return c.loadFromEnvFile()
}
