package main

import (
	"os"

	"gopkg.in/ini.v1"
)

type config struct {
	aws     string
	region  string
	user    string
	port    string
	bastion string
}

func (c config) new() (config, error) {
	cfgFile := os.Getenv("HOME") + "/.gssh"
	cfg, err := ini.Load(cfgFile)

	// get entries
	aws := cfg.Section("default").Key("aws").String()
	region := cfg.Section("default").Key("region").String()
	user := cfg.Section("default").Key("user").String()
	port := cfg.Section("default").Key("port").String()
	bastion := cfg.Section("default").Key("bastion").String()

	return config{aws: aws, region: region, user: user, port: port, bastion: bastion}, err
}
