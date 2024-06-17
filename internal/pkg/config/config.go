package config

import (
	"errors"
	"log"
	"strings"

	"github.com/spf13/viper"

	"github.com/nijeti/cinema-keeper/internal/types"
)

type DiscordCfg struct {
	Token string   `mapstructure:"token"`
	Guild types.ID `mapstructure:"guild"`
}

type DbCfg struct {
	ConnectionString string `mapstructure:"connection_string"`
}

type Config struct {
	Discord DiscordCfg `mapstructure:"discord"`
	DB      DbCfg      `mapstructure:"db"`
}

func ReadConfig() *Config {
	viper.AddConfigPath(".")
	viper.AddConfigPath("./cmd/service/")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil && !errors.As(err, &viper.ConfigFileNotFoundError{}) {
		log.Fatalln("failed to read config file", err)
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("service")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	cfg := &Config{}
	err = viper.Unmarshal(cfg)
	if err != nil {
		log.Fatalln("failed to unmarshal config", err)
	}
	return cfg
}
