package gssh

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/spf13/viper"
)

func init() {
	log.SetFlags(0)

	path := os.Getenv("HOME") + "/.gssh"
	_ = os.Mkdir(path, 0744)

	viper.SetConfigName("config")
	viper.AddConfigPath(path)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Print("AWS profile (default default): ")
		var awsProfile string
		fmt.Scanln(&awsProfile)
		if len(awsProfile) == 0 {
			awsProfile = "default"
		}

		fmt.Print("AWS region (default us-east-1): ")
		var awsRegion string
		fmt.Scanln(&awsRegion)
		if len(awsRegion) == 0 {
			awsRegion = "us-east-1"
		}

		user, err := user.Current()
		if err != nil {
			panic(err)
		}

		fmt.Print("SSH user (default " + user.Name + "): ")
		var sshUser string
		fmt.Scanln(&sshUser)
		if len(sshUser) == 0 {
			sshUser = user.Name
		}

		fmt.Print("SSH port (default 22): ")
		var sshPort int
		fmt.Scanln(&sshPort)
		if sshPort == 0 {
			sshPort = 22
		}

		fmt.Print("SSH bastion (default empty): ")
		var sshBastion string
		fmt.Scanln(&sshBastion)

		viper.Set("aws.profile", awsProfile)
		viper.Set("aws.region", awsRegion)
		viper.Set("ssh.user", sshUser)
		viper.Set("ssh.port", sshPort)
		viper.Set("ssh.bastion", sshBastion)

		if err := viper.WriteConfigAs(path + "/config.yaml"); err != nil {
			log.Fatal(err)
		}
	}
}
