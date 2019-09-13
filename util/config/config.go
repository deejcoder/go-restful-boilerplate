package config

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config wraps the config file as a struct
type Config struct {
	Version      string
	LogLevel     logrus.Level
	AppName      string
	AppShortName string
	API          apiConfig
	Database     databaseConfig
}

type apiConfig struct {
	Port           int
	AllowedMethods []string
	AllowedHeaders []string
	AllowedOrigins []string
}

type databaseConfig struct {
	Host string
	Port int
}

// init serializes YAML into a Config struct
func (cfg *Config) init() {
	cfg.Version = viper.GetString("version")
	cfg.setLogLevel(viper.GetString("log_level"))
	cfg.AppName = viper.GetString("app_name")
	cfg.AppShortName = viper.GetString("app_short_name")
	cfg.API.Port = viper.GetInt("api.port")
	cfg.API.AllowedMethods = viper.GetStringSlice("api.allowed_methods")
	cfg.API.AllowedHeaders = viper.GetStringSlice("api.allowed_headers")
	cfg.API.AllowedOrigins = viper.GetStringSlice("api.allowed_origins")
	cfg.Database.Host = viper.GetString("database.host")
	cfg.Database.Port = viper.GetInt("database.port")
}

// GetConfig loads config data into a Config struct
func GetConfig() *Config {
	cfg := new(Config)
	cfg.init()

	return cfg
}

// InitConfig sets up the config file
func InitConfig() (*Config, error) {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	cfg := GetConfig()
	log.SetLevel(cfg.LogLevel)

	return cfg, nil
}

func (cfg *Config) setLogLevel(loglevel string) {
	switch loglevel {
	case "info":
		cfg.LogLevel = logrus.InfoLevel
	case "warn":
		cfg.LogLevel = logrus.WarnLevel
	case "fatal":
		cfg.LogLevel = logrus.FatalLevel
	}
}