package config

type Config struct {
	HTTPPort string

	PostgresHost     string
	PostgresUser     string
	PostgresDatabase string
	PostgresPassword string
	PostgresPort     string
	PostgresMaxConn  int32

	RedisAddr     string
	RedisPassword string
	RedisDB       int

	AuthSecretKey string
}

func Load() Config {

	var cfg Config

	cfg.HTTPPort = ":9090"

	cfg.PostgresHost = "localhost"
	cfg.PostgresUser = "postgres"
	cfg.PostgresDatabase = "catalog"
	cfg.PostgresPassword = "0003"
	cfg.PostgresPort = "3003"
	cfg.PostgresMaxConn = 30

	cfg.RedisAddr = "localhost:6379"
	cfg.RedisDB = 0

	cfg.AuthSecretKey = "9K+WgNTglA44Hg=="

	return cfg
}
