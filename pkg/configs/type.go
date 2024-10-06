package configs

type Config struct {
	AppName        string `env:"APP_NAME"`
	AppPort        string `env:"APP_PORT"`
	AppEnv         string `env:"APP_ENV"`
	DbName         string `env:"DB_NAME"`
	DbUsername     string `env:"DB_USERNAME"`
	DbPassword     string `env:"DB_PASSWORD"`
	DbHost         string `env:"DB_HOST"`
	DbPort         int    `env:"DB_PORT"`
	OTLPEndpoint   string `env:"OTEL_RECEIVER_OTLP_ENDPOINT"`
	RedisDB        string `env:"REDIS_DB"`
	RedisHost      string `env:"REDIS_HOST"`
	RedisPort      string `env:"REDIS_PORT"`
	RedisPassword  string `env:"REDIS_PASSWORD"`
	RedisAppConfig string `env:"REDIS_APP_CONFIG"`
	AMQPURL        string `env:"AMQP_URL"`
	RidesAPIKey    string `env:"RIDES_API_KEY"`
}

type ConfigLoader struct {
	Env           string
	ConsulAddress string
}
