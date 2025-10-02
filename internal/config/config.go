package config

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Env               string        `mapstructure:"env"`
	Port              string        `mapstructure:"port"`
	UIDevMode         bool          `mapstructure:"ui_dev_mode"`
	UIDevServerURL    string        `mapstructure:"ui_dev_server_url"`
	ReadTimeout       time.Duration `mapstructure:"read_timeout"`
	WriteTimeout      time.Duration `mapstructure:"write_timeout"`
	ReadHeaderTimeout time.Duration `mapstructure:"read_header_timeout"`
}

// Load reads configuration from .env file and environment variables.
// Environment variables take precedence over .env file values.
func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	// Set defaults
	viper.SetDefault("env", "production")
	viper.SetDefault("port", "8080")
	viper.SetDefault("ui_dev_mode", false)
	viper.SetDefault("ui_dev_server_url", "http://localhost:5173")
	viper.SetDefault("read_timeout", 5*time.Minute)
	viper.SetDefault("write_timeout", 5*time.Minute)
	viper.SetDefault("read_header_timeout", 1*time.Minute)

	// Read config file (optional, won't fail if missing)
	_ = viper.ReadInConfig()

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return &cfg, nil
}

// Validate checks if configuration values are valid.
func (c *Config) Validate() error {
	if c.Port == "" {
		return errors.New("port cannot be empty")
	}

	if c.UIDevMode && c.UIDevServerURL == "" {
		return errors.New("ui_dev_server_url cannot be empty when ui_dev_mode is enabled")
	}

	if c.ReadTimeout < 0 {
		return errors.New("read_timeout cannot be negative")
	}

	if c.WriteTimeout < 0 {
		return errors.New("write_timeout cannot be negative")
	}

	if c.ReadHeaderTimeout < 0 {
		return errors.New("read_header_timeout cannot be negative")
	}

	return nil
}
