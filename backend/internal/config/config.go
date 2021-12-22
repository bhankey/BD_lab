package config

import (
	"finance/pkg/logger"
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

// Config struct that depends configuration of App
type Config struct {
	Server struct {
		Port         string `yaml:"port" env:"PORT" env-default:"8080"`
		WriteTimeout int    `yaml:"write_timeout" env-default:"15"`
		ReadTimeout  int    `yaml:"read_timeout" env-default:"15"`
		IdleTimeout  int    `yaml:"idle_timeout" env-default:"30"`
	}
	Postgres struct {
		Host     string `yaml:"host" env:"PG_HOST" env-default:"localhost"`
		Port     string `yaml:"port" env:"PG_PORT" env-default:"5432"`
		User     string `yaml:"user" env:"PG_USER" env-default:"postgres"`
		Password string `yaml:"password" env:"PG_PASSWORD" env-default:"postgres"`
		DBName   string `yaml:"db_name" env:"PG_NAME" env-default:"finance"`
	}
	Redis struct {
		Host     string `yaml:"host" env:"RD_HOST" env-default:"localhost"`
		Port     string `yaml:"port" env:"RD_PORT" env-default:"6379"`
		Password string `yaml:"password" env:"RD_PASSWORD" env-default:"redis"`
	}
}

var instance *Config
var once sync.Once

// GetConfig return pointer to config. Config is singleton
func GetConfig(path string) *Config {
	once.Do(func() {
		l := logger.GetLogger()
		l.Infoln("reading server config file")
		instance = &Config{}
		if path == "" {
			path = "./config/config.yaml"
		}
		if err := cleanenv.ReadConfig(path, instance); err != nil {
			l.Fatalln(err)
		}
	})
	return instance
}
