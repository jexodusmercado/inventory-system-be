package conf

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port               string `mapstructure:"PORT"`
	DbHost             string `mapstructure:"DB_HOST"`
}

func LoadConfig() (conf Config) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalln("Error loading config file")
		return Config{}
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		log.Fatalln("Unable to decode into struct")
		return Config{}
	}

	return conf
}