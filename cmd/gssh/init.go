package main

import (
	"fmt"
	"os"
)

// configure gssh
func init() {
	cfgFile := os.Getenv("HOME") + "/.gssh"

	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		fmt.Print("AWS profile (Default: default): ")
		var aws string
		fmt.Scanln(&aws)
		if len(aws) == 0 {
			aws = "default"
		}

		fmt.Print("AWS region (Default: us-east-1): ")
		var region string
		fmt.Scanln(&region)
		if len(region) == 0 {
			region = "us-east-1"
		}

		fmt.Print("SSH user: ")
		var user string
		fmt.Scanln(&user)

		fmt.Print("SSH port: ")
		var port string
		fmt.Scanln(&port)

		fmt.Print("SSH bastion (Default: empty): ")
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
