package util

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	Environment string `mapstructure:"ENVIRONMENT"`
	DBSource    string `mapstructure:"DB_SOURCE"`
	// MigrationURL         string        `mapstructure:"MIGRATION_URL"`
	// RedisAddress         string        `mapstructure:"REDIS_ADDRESS"`
	HTTPServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress string `mapstructure:"GRPC_SERVER_ADDRESS"`
	// TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	// AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	// RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	// EmailSenderName      string        `mapstructure:"EMAIL_SENDER_NAME"`
	// EmailSenderAddress   string        `mapstructure:"EMAIL_SENDER_ADDRESS"`
	// EmailSenderPassword  string        `mapstructure:"EMAIL_SENDER_PASSWORD"`
	DB_CONTAINER_NAME string `mapstructure:"DB_CONTAINER_NAME"`
	DB_USER           string `mapstructure:"DB_USER"`
	DB_NAME           string `mapstructure:"DB_NAME"`
	DB_PORT           string `mapstructure:"DB_PORT"`
	NETWORK_NAME      string `mapstructure:"NETWORK_NAME"`
	DB_HOST           string `mapstructure:"DB_HOST"`
	DB_PASSWORD       string `mapstructure:"DB_PASSWORD"`
	DRIVER            string `mapstructure:"DRIVER"`
}

var config Config

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (Config, error) {
	if _, err := os.Stat(fmt.Sprintf("%s/app.env", path)); err == nil {
		viper.AddConfigPath(path)
		viper.SetConfigName("app")
		viper.SetConfigType("env")

		if err := viper.ReadInConfig(); err != nil {
			return config, fmt.Errorf("read config error %v", err)
		}
	}

	viper.AutomaticEnv()

	// Map environment variable names to viper keys
	viper.SetEnvPrefix("") // This allows reading env vars without a prefix
	err := viper.BindEnv("DB_USER")
	if err != nil {
		return config, fmt.Errorf("error binding DB_USER environment variable: %v", err)
	}
	err = viper.BindEnv("DB_PASSWORD")
	if err != nil {
		return config, fmt.Errorf("error binding  environment variable: %v", err)
	}
	err = viper.BindEnv("DB_NAME")
	if err != nil {
		return config, fmt.Errorf("error binding environment variable: %v", err)
	}
	err = viper.BindEnv("DB_PORT")
	if err != nil {
		return config, fmt.Errorf("error binding  environment variable: %v", err)
	}
	err = viper.BindEnv("DB_HOST")
	if err != nil {
		return config, fmt.Errorf("error binding  environment variable: %v", err)
	}
	err = viper.BindEnv("DB_CONTAINER_NAME")
	if err != nil {
		return config, fmt.Errorf("error binding  environment variable: %v", err)
	}
	err = viper.BindEnv("HTTP_SERVER_ADDRESS")
	if err != nil {
		return config, fmt.Errorf("error binding  environment variable: %v", err)
	}
	err = viper.BindEnv("NETWORK_NAME")
	if err != nil {
		return config, fmt.Errorf("error binding  environment variable: %v", err)
	}
	err = viper.BindEnv("DRIVER")
	if err != nil {
		return config, fmt.Errorf("error binding  environment variable: %v", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}

	if config.DB_USER != "" && config.DB_HOST != "" && config.DB_PORT != "" && config.DB_NAME != "" {
		config.DBSource = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME)
	}
	log.Printf("source%v", config.DBSource)
	return config, nil
}
