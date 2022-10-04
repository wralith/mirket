package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	*AppConfig
}

type AppConfig struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Logger   LoggerConfig
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
}

type LoggerConfig struct {
}

func NewConfig() *Config {
	appConfig := &AppConfig{}
	appConfig.readAppConfig()

	return &Config{AppConfig: appConfig}
}

func (c *AppConfig) readAppConfig() {
	env := activeEnv()

	log.Println("ACTIVE_PROFILE: ", env)
	v := viper.New()

	v.SetTypeByDefaultValue(true)
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")

	err := v.ReadInConfig()

	if err != nil {
		panic("Unable to load app config, terminating: " + err.Error())
	}

	sub := v.Sub(env)
	err = sub.Unmarshal(c)
	if err != nil {
		panic("Unable to deserialize app config, terminating: " + err.Error())
	}
}

func activeEnv() string {
	env, found := os.LookupEnv("ACTIVE_PROFILE")

	if !found {
		env = "local"
	}

	return env
}
