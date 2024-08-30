package config

import (
	"db_cp_6/pkg/logger"
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type Config struct {
	HTTPServer `yaml:"http_server"`
	Member     Postgres `yaml:"memberpostgres"`
	Leader     Postgres `yaml:"leaderpostgres"`
	Admin      Postgres `yaml:"adminpostgres"`
	Test       Postgres `yaml:"testpostgres"`
}

type HTTPServer struct {
	Host string `yaml:"host"`
	Port string `yaml:"port" default:"8080"`
}

type Postgres struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port" default:"5432"`
	Database string `yaml:"dbname"`
}

var instance *Config
var once sync.Once

func GetConfig(logger *logger.Logger) *Config {
	once.Do(func() {
		logger.Info("read application configuration")
		instance = &Config{}
		if err := cleanenv.ReadConfig("./config/config.yaml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
