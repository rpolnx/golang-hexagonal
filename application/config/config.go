package config

import (
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Configuration struct {
	Database
	App
}

type Database struct {
	Mongo
	Redis
}

type Mongo struct {
	Uri     string
	Db      string
	Timeout int
}

type Redis struct {
	Host        string
	Username    string
	Password    string
	Db          int
	Timeout     int
	Ttl         int
	PoolSize    int
	IdleTimeout int
}

type App struct {
	Timeout int
	Port    int
}

func LoadConfig() (config *Configuration, e error) {

	if os.Getenv("ENV") == "PRODUCTION" {
		viper.SetConfigName("config")
	} else {
		viper.SetConfigName("application")
	}

	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "config.LoadConfig")
	}

	err := viper.Unmarshal(&config)

	if err != nil {
		return nil, errors.Wrap(err, "config.LoadConfig")
	}

	return config, nil
}
