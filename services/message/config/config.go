package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func LoadConfig(config *Config) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./services/auth/config")
	viper.AutomaticEnv()

	// Read in the configuration file
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode config into struct: %v", err)
	}
}
