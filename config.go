package main

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ConnectionString string `mapstructure:"connection_string"`
	Port             string `mapstructure:"port"`
}

var AppConfig *Config

func loadAppConfig() {
	log.Println("Loading Server Configurations...")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}

}
