package config

import (
	"github.com/spf13/viper"
	"log"
)

var config *Config

type Config struct {
	// Define your configuration fields here
	Woocommerce struct {
		ConsumerKey    string
		ConsumerSecret string
		WebsiteUrl     string
	}
}

// LoadConfig loads the configuration from a file
func LoadConfig(filePath string) error {
	viper.SetConfigFile(filePath)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
		return err
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct: %s", err)
		return err
	}

	return nil
}

func GetConfig() *Config {
	return config
}
