package main

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	Port       string `mapstructure:"port"`
	Connstring string `mapstructure:"connection"`
}

var AppConfig *Config

func loadAppConfig() {
	log.Info().Msg("Loading server configurations...")

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		log.Error().Err(err).Msg("Error reading config...")
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Error().Err(err).Msg("Error unmarshalling config...")
	}
}

func init() {}

func main() {
}
