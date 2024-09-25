package config

import "time"

type RedisConfig struct {
	RedisCacheHost string `env:"REDIS_CACHE_HOST" envDefault:"127.0.0.1"`
	RedisCachePort int    `env:"REDIS_CACHE_PORT" envDefault:"6379"`

	RedisMaxOpenConns    int           `env:"REDIS_MAX_OPEN_CONNS" envDefault:"100"`
	RedisMaxIdleConns    int           `env:"REDIS_MAX_IDLE_CONNS" envDefault:"20"`
	RedisConnMaxIdleTime time.Duration `env:"REDIS_CONN_MAX_IDLE_TIME" envDefault:"30s"`
	RedisConnMaxLifetime time.Duration `env:"REDIS_CONN_MAX_LIFETIME" envDefault:""`

	RedisName   string        `env:"REDIS_Name" envDefault:""`
	RedisMaxTTL time.Duration `env:"REDIS_MAX_TTL" envDefault:"720h"`
}
