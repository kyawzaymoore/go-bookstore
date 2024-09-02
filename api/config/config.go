package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

// SecretKey is the key used for signing JWT tokens
//var SecretKey = []byte("fido-golang-key")

type Config struct {
    Server   ServerConfig   `mapstructure:"server"`
    Database DatabaseConfig `mapstructure:"database"`
    App      AppKey      `mapstructure:"app"`
}

type ServerConfig struct {
    Port int    `mapstructure:"port"`
    Host string `mapstructure:"host"`
}

type DatabaseConfig struct {
    Driver   string `mapstructure:"driver"`
    Host     string `mapstructure:"host"`
    Port     int    `mapstructure:"port"`
    User     string `mapstructure:"user"`
    Password string `mapstructure:"password"`
    DBName   string `mapstructure:"dbname"`
    ConnectionString   string `mapstructure:"connectionString"`
}

type AppKey struct {
    SecretKey string `mapstructure:"secretKey"`
}

var AppConfig *Config

func LoadConfig() {
    configPath := os.Getenv("CONFIG_PATH")  // Get the config path from the environment variable

    if configPath == "" {
        configPath = "config/config.dev.json"
    }

    viper.SetConfigFile(configPath)
    viper.AutomaticEnv()

    err := viper.ReadInConfig()
    if err != nil {
        log.Fatalf("Error reading config file, %s", err)
    }

    AppConfig = &Config{}
    err = viper.Unmarshal(AppConfig)
    if err != nil {
        log.Fatalf("Unable to decode into struct, %v", err)
    }
}