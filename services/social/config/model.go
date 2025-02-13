package config

import (
	"time"
)

type ServerConfig struct {
	Host         string        `mapstructure:"host"`
	Port         int           `mapstructure:"port"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`  // in seconds
	WriteTimeout time.Duration `mapstructure:"write_timeout"` // in seconds
	IdleTimeout  time.Duration `mapstructure:"idle_timeout"`  // in seconds
	EnableCORS   bool          `mapstructure:"enable_cors"`
}

type PostgresConfig struct {
	Host               string `mapstructure:"host"`
	Port               int    `mapstructure:"port"`
	User               string `mapstructure:"user"`
	Password           string `mapstructure:"password"`
	DBName             string `mapstructure:"db_name"`
	SSLMode            string `mapstructure:"ssl_mode"`
	MaxConnections     int    `mapstructure:"max_connections"`
	MaxIdleConnections int    `mapstructure:"max_idle_connections"`
}

type RedisConfig struct {
	Host          string `mapstructure:"host"`
	Port          int    `mapstructure:"port"`
	Password      string `mapstructure:"password"`
	DB            int    `mapstructure:"db"`
	SessionExpiry int    `mapstructure:"session_expiry"` // in seconds
}

type LoggingConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
	Output string `mapstructure:"output"`
}

type FeaturesConfig struct {
	EnableRegistration       bool `mapstructure:"enable_registration"`
	EnableEmailNotifications bool `mapstructure:"enable_email_notifications"`
	Enable2FA                bool `mapstructure:"enable_2fa"`
}

type RateLimitConfig struct {
	Enabled       bool `mapstructure:"enabled"`
	MaxRequests   int  `mapstructure:"max_requests"`
	WindowSeconds int  `mapstructure:"window_seconds"`
}

type CORSConfig struct {
	AllowedOrigins []string `mapstructure:"allowed_origins"`
	AllowedMethods []string `mapstructure:"allowed_methods"`
	AllowedHeaders []string `mapstructure:"allowed_headers"`
}

type APIConfig struct {
	RateLimit RateLimitConfig `mapstructure:"rate_limit"`
	CORS      CORSConfig      `mapstructure:"cors"`
}

type Config struct {
	Server   ServerConfig `mapstructure:"server"`
	Database struct {
		Postgres PostgresConfig `mapstructure:"postgres"`
	} `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	Logging  LoggingConfig  `mapstructure:"logging"`
	Features FeaturesConfig `mapstructure:"features"`
	API      APIConfig      `mapstructure:"api"`
}
