package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config ...
type Config struct{}

const (
	DefaultConfigPath = ""
)

func NewConfig(path string) *Config {
	c := Config{}

	return &c
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if DefaultConfigPath != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(DefaultConfigPath)
	}

	viper.SetConfigName(".miner") // name of config file (without extension)
	viper.AddConfigPath(".")      // adding current directory
	viper.AutomaticEnv()          // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
