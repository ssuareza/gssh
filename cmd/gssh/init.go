package main

import (
	"fmt"
	"os"
)

// configure gssh
func init() {
	cfgFile := os.Getenv("HOME") + "/.gssh"

	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		fmt.Print("AWS profile (default: default): ")
		var aws string
		fmt.Scanln(&aws)
		if len(aws) == 0 {
			aws = "default"
		}

		fmt.Print("AWS region (default: us-east-1): ")
		var region string
		fmt.Scanln(&region)
		if len(region) == 0 {
			region = "us-east-1"
		}

		fmt.Print("SSH user: ")
		var user string
		fmt.Scanln(&user)

		fmt.Print("SSH port: (default: 22)")
		var port string
		fmt.Scanln(&port)
		if len(port) == 0 {
			region = "22"
		}

		fmt.Print("SSH bastion (default: empty): ")
		var bastion string
		fmt.Scanln(&bastion)

		content := `[default]
aws = ` + aws + `
region = ` + region + `
user = ` + user + `
port = ` + port + `
bastion = ` + bastion

		f, _ := os.Create(cfgFile)
		f.WriteString(content + "\n")
		f.Close()
	}
}
