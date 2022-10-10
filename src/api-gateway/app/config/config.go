package config

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Logger   LoggerConfig
	Services ServicesConfig
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type LoggerConfig struct {
	Level  string `yaml:"level"`
	Pretty bool   `yaml:"pretty"`
}

type ServicesConfig struct {
	User string `yaml:"user"`
}

func NewConfig() *Config {
	c := &Config{}
	c.readConfig()
	return c
}

func (c *Config) readConfig() {
	env := activeEnv()
	v := viper.New()

	v.SetTypeByDefaultValue(true)
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./app/config")

	err := v.ReadInConfig()
	if err != nil {
		log.Panic().Err(err).Msg("Unable to load config")
	}

	sub := v.Sub(env)
	err = sub.Unmarshal(c)
	if err != nil {
		log.Panic().Err(err).Msg("Unable to unmarshal config")
	}
}

func activeEnv() string {
	env, found := os.LookupEnv("ACTIVE_PROFILE")
	if !found {
		env = "local"
	}
	return env
}
