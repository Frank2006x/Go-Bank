package util




type Config struct {
	DBSource string `env:"DB_SOURCE"`
	ServerAddress string `env:"SERVER_ADDRESS"`
}

