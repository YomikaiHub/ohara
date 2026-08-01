package config

import "github.com/YomikaiHub/ohara/api/internal/env"

func Load(version string) *Config {
	return &Config{
		App: AppConfig{
			Addr:    env.MustGet("ADDR"),
			Env:     env.MustGet("ENV"),
			Version: version,
		},

		DB: DBConfig{
			URL:          env.MustGet("DB_URL"),
			MaxOpenConns: env.MustGetInt("DB_MAX_OPEN_CONNS"),
			MaxIdleConns: env.MustGetInt("DB_MAX_IDLE_CONNS"),
			MaxIdleTime:  env.MustGetDuration("DB_MAX_IDLE_TIME"),
		},

		JWT: JWTConfig{
			Key:             env.MustGet("JWT_SECRET"),
			AccessTokenTTL:  env.MustGetDuration("JWT_ACCESS_TOKEN_TTL"),
			RefreshTokenTTL: env.MustGetDuration("JWT_REFRESH_TOKEN_TTL"),
		},
	}
}
