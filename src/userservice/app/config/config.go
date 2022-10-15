package config

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
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
	Level  string `yaml:"level"`
	Pretty bool   `yaml:"pretty"`
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

	port := os.Getenv("PORT")
	if port != "" {
		c.Server.Port = port
	}
}

func activeEnv() string {
	env, found := os.LookupEnv("ACTIVE_PROFILE")
	if !found {
		env = "local"
	}
	return env
}
