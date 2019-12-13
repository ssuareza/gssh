package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/gosuri/uitable"
	"github.com/ssuareza/gssh"
)

func print(i []map[string]string) {
	table := uitable.New()
	table.MaxColWidth = 50
	table.AddRow("InstanceID", "Name", "PrivateIP", "PublicIP")
	for _, instance := range i {
		table.AddRow(instance["instance-id"], instance["tag:Name"], instance["private-ip"], instance["public-ip"])
	}
	fmt.Printf("%s\n\n", table)
}

func getIP(id string, i []map[string]string, iptype string) (string, error) {
	for _, instance := range i {
		if instance["instance-id"] == id {
			if iptype == "public" {
				return instance["public-ip"], nil
			}
			return instance["private-ip"], nil
		}
	}
	return "", errors.New("IP not found")
}

func main() {
	// get config
	c, err := gssh.GetConfig()
	if err != nil {
		log.Panic(err)
	}

	// get instances
	profiles := strings.Split(c.AWS.Profile, ",")
	var instances []*ec2.DescribeInstancesOutput
	for k := range profiles {
		svc, err := gssh.NewService(profiles[k], c.AWS.Region)
		if err != nil {
			log.Fatal(err)
		}

		list, err := gssh.Get(svc)
		if err != nil {
			log.Fatal(err)
		}
		instances = append(instances, list)
	}

	i := gssh.Filter(instances)
	if err != nil {
		log.Panic(err)
	}
	print(i)

	// select instance
	fmt.Print("Select InstanceID: ")
	var instanceID string
	_, err = fmt.Scanln(&instanceID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// get ip
	iptype := "private"
	if len(c.SSH.Bastion) == 0 {
		iptype = "public"
	}

	ip, err := (getIP(instanceID, i, iptype))
	if err != nil {
		log.Fatal(err)
	}

	// and connect
	gssh.Shell(ip, c)
}
