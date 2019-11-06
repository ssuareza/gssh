package gssh

import (
	"os"

	"gopkg.in/ini.v1"
)

// Config holds gssh configuration
type Config struct {
	AWS     string
	Region  string
	User    string
	Port    string
	Bastion string
}

// ReadConfig reads gssh configuration
func ReadConfig() (Config, error) {
	cfgFile := os.Getenv("HOME") + "/.gssh"
	cfg, err := ini.Load(cfgFile)

	// get entries
	aws := cfg.Section("default").Key("aws").String()
	region := cfg.Section("default").Key("region").String()
	user := cfg.Section("default").Key("user").String()
	port := cfg.Section("default").Key("port").String()
	bastion := cfg.Section("default").Key("bastion").String()

	return Config{AWS: aws, Region: region, User: user, Port: port, Bastion: bastion}, err
}
