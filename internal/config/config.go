package config

import "github.com/caarlos0/env/v11"

type HTTP struct {
	Addr string `env:"APP_ADDR,default=:8080"`
}
type DB struct {
	DSN string `env:"DB_DSN,required"`
}
type Minio struct {
	Endpoint  string `env:"MINIO_ENDPOINT,default=localhost:9000"`
	AccessKey string `env:"MINIO_ACCESS_KEY,default=minio"`
	SecretKey string `env:"MINIO_SECRET_KEY,default=minio12345"`
	Bucket    string `env:"MINIO_BUCKET,default=attachments"`
	UseSSL    bool   `env:"MINIO_USE_SSL,default=false"`
}
type Config struct {
	HTTP  HTTP
	DB    DB
	Minio Minio
}

func Load() (Config, error) {
	var c Config
	return c, env.Parse(&c)
}
