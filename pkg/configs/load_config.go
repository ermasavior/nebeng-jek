package configs

import (
	"fmt"
	"log"

	consulPkg "nebeng-jek/pkg/consul"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

func (c *ConfigLoader) loadFromEnvFile() *Config {
	var err error
	if c.Env == "" {
		err = godotenv.Load("./configs/.env")
	} else {
		err = godotenv.Load(fmt.Sprintf("./configs/%s.env", c.Env))
	}

	// log file is optional
	if err != nil {
		log.Println("error loading config from file ", err)
	}

	cfg := new(Config)
	if err := env.Parse(cfg); err != nil {
		log.Fatal("error parsing config from file ", err)
	}
	return cfg
}

func (c *ConfigLoader) loadFromConsul() (*Config, error) {
	kv, err := consulPkg.NewConsulKVClient(c.ConsulAddress)
	if err != nil {
		return nil, err
	}

	appNameBytes, err := kv.GetKeyValue("APP_NAME", nil)
	if err != nil {
		return nil, err
	}
	appPortBytes, err := kv.GetKeyValue("APP_PORT", nil)
	if err != nil {
		return nil, err
	}
	dbUsernameBytes, err := kv.GetKeyValue("DB_USERNAME", nil)
	if err != nil {
		return nil, err
	}
	dbPasswordBytes, err := kv.GetKeyValue("DB_PASSWORD", nil)
	if err != nil {
		return nil, err
	}
	dbHostBytes, err := kv.GetKeyValue("DB_HOST", nil)
	if err != nil {
		return nil, err
	}
	redisDbBytes, err := kv.GetKeyValue("REDIS_DB", nil)
	if err != nil {
		return nil, err
	}
	redisHostBytes, err := kv.GetKeyValue("REDIS_HOST", nil)
	if err != nil {
		return nil, err
	}
	redisPortBytes, err := kv.GetKeyValue("REDIS_PORT", nil)
	if err != nil {
		return nil, err
	}
	redisPasswordBytes, err := kv.GetKeyValue("REDIS_PASSWORD", nil)
	if err != nil {
		return nil, err
	}
	redisAppConfigBytes, err := kv.GetKeyValue("REDIS_APP_CONFIG", nil)
	if err != nil {
		return nil, err
	}

	return &Config{
		AppName:        string(appNameBytes),
		AppPort:        string(appPortBytes),
		DbUsername:     string(dbUsernameBytes),
		DbPassword:     string(dbPasswordBytes),
		DbHost:         string(dbHostBytes),
		RedisDB:        string(redisDbBytes),
		RedisHost:      string(redisHostBytes),
		RedisPort:      string(redisPortBytes),
		RedisPassword:  string(redisPasswordBytes),
		RedisAppConfig: string(redisAppConfigBytes),
	}, nil
}
