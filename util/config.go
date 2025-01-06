package util

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	Environment          string        `mapstructure:"ENVIRONMENT"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	MigrationURL         string        `mapstructure:"MIGRATION_URL"`
	RedisAddress         string        `mapstructure:"REDIS_ADDRESS"`
	HTTPServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress    string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	EmailSenderName      string        `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress   string        `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword  string        `mapstructure:"EMAIL_SENDER_PASSWORD"`
	DB_CONTAINER_NAME    string        `mapstructure:"DB_CONTAINER_NAME"`
	DB_USER              string        `mapstructure:"DB_USER"`
	DB_NAME              string        `mapstructure:"DB_NAME"`
	DB_PORT              string        `mapstructure:"DB_PORT"`
	NETWORK_NAME         string        `mapstructure:"NETWORK_NAME"`
	DB_HOST              string        `mapstructure:"DB_HOST"`
	DB_PASSWORD          string        `mapstructure:"DB_PASSWORD"`
	DRIVER               string        `mapstructure:"DRIVER"`
}

var config Config

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (Config, error) {

	env := viper.GetString("APP_ENV")
	if env == "" {
		env = "dev" // Default to development environment
	}
	configFileName := fmt.Sprintf("app.%s", env)
	path = fmt.Sprintf("%s/config", path)
	viper.AddConfigPath(path)
	viper.SetConfigName(configFileName)
	viper.SetConfigType("env")
	log.Printf("env=%s", env)
	if err := viper.ReadInConfig(); err != nil {
		return config, fmt.Errorf("read config error %v", err)
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return config, fmt.Errorf("read config error %v", err)
	}
	err := viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}

	if config.DBSource == "" {
		if config.DB_USER != "" && config.DB_HOST != "" && config.DB_PORT != "" && config.DB_NAME != "" {
			config.DBSource = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME)
			log.Printf("config.DBSource=%s", config.DBSource)
		} else {
			return config, fmt.Errorf("DB Source missing %v", err)
		}
	}
	fmt.Printf("DbSource:%s", config.DBSource)
	return config, nil
}
