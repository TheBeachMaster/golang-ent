package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DB     PostgresConfig
	Cache  RedisConfig
	Server ServerConfig
	Logger LoggerConfig
}

type RedisConfig struct {
	RedisUser string
	RedisPass string
	RedisHost string
	RedisPort string
}

type PostgresConfig struct {
	DBUser string
	DBPass string
	DBHost string
	DBPort string
	DBName string
}

type ServerConfig struct {
	ReadTimeout  uint8
	AppName      string
	ServerHeader string
	Prefork      bool
	Port         int
}

type LoggerConfig struct {
	LogFile           string
	Timezone          string
	LoggerMaxFileSize uint8
	CompressLogFile   bool
	UseLocalTime      bool
}

func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

// Parse config file
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}
