package config

import "time"

type Config struct {
	App AppConfig
	DB  DBConfig
	JWT JWTConfig
}

type AppConfig struct {
	Env     string
	Addr    string
	Version string
}

type JWTConfig struct {
	Key             string
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

type DBConfig struct {
	URL          string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  time.Duration
}
