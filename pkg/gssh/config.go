package gssh

import (
	"os"

	"github.com/spf13/viper"
)

// AWS holds aws configuration
type AWS struct {
	Profile string
	Region  string
}

// SSH holds ssh configuration
type SSH struct {
	User    string
	Port    int
	Bastion string
}

// Config holds all configuration
type Config struct {
	AWS *AWS
	SSH *SSH
}

// GetConfig reads gssh configuration
func GetConfig() (*Config, error) {
	path := os.Getenv("HOME") + "/.gssh"
	viper.SetConfigName("config")
	viper.AddConfigPath(path)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return &Config{}, err
	}

	return &Config{
		&AWS{
			Profile: viper.GetString("aws.profile"),
			Region:  viper.GetString("aws.region"),
		},
		&SSH{
			User:    viper.GetString("ssh.user"),
			Port:    viper.GetInt("ssh.port"),
			Bastion: viper.GetString("ssh.bastion"),
		},
	}, nil
}
